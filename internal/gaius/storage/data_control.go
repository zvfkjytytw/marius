package gaiusstorage

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/zvfkjytytw/marius/proto/mulus/api/v1"
)

const fileNameTemplate = `file_%d`

func (s *Storage) GetFile(ctx context.Context, fileID int32) ([]byte, error) {
	mulus, err := s.tesser.GetMulus(fileID)
	if err != nil {
		return nil, fmt.Errorf("fail get mulus for file %d: %w", fileID, err)
	}

	var dataSlice [dataParts][]byte
	fileName := fmt.Sprintf(fileNameTemplate, fileID)

	var wg sync.WaitGroup
	wg.Add(dataParts)
	for i, mus := range mulus {
		go func(musAddress string, part int) {
			defer wg.Done()
			data, err := s.getData(ctx, fileName, musAddress)
			if err != nil {
				s.logger.Sugar().Errorf("fail get data from mus %s: %v", musAddress, err)
			}
			dataSlice[part] = data
		}(mus, i)
	}
	wg.Wait()

	var fileData []byte
	for _, data := range dataSlice {
		fileData = append(fileData, data...)
	}

	return fileData, nil
}

func (s *Storage) getData(ctx context.Context, fileName, musAddress string) ([]byte, error) {
	resp, err := s.castra[musAddress].client.GetData(
		ctx,
		&api.GetRequest{
			Name: fileName,
		},
	)

	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *Storage) SaveFile(ctx context.Context, fileID int32, data []byte) error {
	mulusPrima, err := s.selectMulus()
	if err != nil {
		return fmt.Errorf("fail to select mulus: %w", err)
	}

	var mulusExitus [dataParts]string
	indexStep := len(data) / dataParts
	fileName := fmt.Sprintf(fileNameTemplate, fileID)

	var wg sync.WaitGroup
	wg.Add(dataParts)
	for i := 0; i < dataParts; i++ {
		firstIndex := i * indexStep
		lastIndex := (i + 1) * indexStep
		if i == dataParts-1 {
			lastIndex = len(data)
		}
		go func(data []byte, part int) {
			musAddress := mulusPrima[part].address
			defer wg.Done()
			if err := s.saveData(ctx, fileName, musAddress, data); err != nil {
				s.logger.Sugar().Errorf("fail save data on mus %s: %v", musAddress, err)
			}
			mulusPrima[part].Lock()
			mulusPrima[part].sarcina++
			mulusPrima[part].Unlock()
			mulusExitus[part] = musAddress
		}(data[firstIndex:lastIndex], i)
	}
	wg.Wait()

	sort.Slice(s.cohors, func(i, j int) bool { return s.cohors[i].sarcina < s.cohors[j].sarcina })
	if err := s.tesser.SetMulus(fileID, mulusExitus); err != nil {
		return fmt.Errorf("fail set mulus for file %d: %w", fileID, err)
	}

	return nil
}

func (s *Storage) saveData(ctx context.Context, fileName, musAddress string, data []byte) error {
	resp, err := s.castra[musAddress].client.SaveData(
		ctx,
		&api.SaveRequest{
			Name: fileName,
			Data: data,
		},
	)

	if err != nil {
		return err
	}

	if resp.Name != fileName {
		return fmt.Errorf("wrong part name %s", resp.Name)
	}

	return nil
}

func (s *Storage) UpdateFile(ctx context.Context, fileID int32, data []byte) error {
	if err := s.DeleteFile(ctx, fileID); err != nil {
		return err
	}

	return s.SaveFile(ctx, fileID, data)
}

func (s *Storage) DeleteFile(ctx context.Context, fileID int32) error {
	mulus, err := s.tesser.GetMulus(fileID)
	if err != nil {
		return fmt.Errorf("fail to get mulus: %w", err)
	}

	fileName := fmt.Sprintf(fileNameTemplate, fileID)

	var wg sync.WaitGroup
	wg.Add(dataParts)
	for _, mus := range mulus {
		go func(musAddress string) {
			defer wg.Done()
			if err := s.deleteData(ctx, fileName, musAddress); err != nil {
				s.logger.Sugar().Errorf("fail to delete file %s from mus %s: %v", fileName, mus, err)
			}
		}(mus)
	}
	wg.Wait()

	if err := s.tesser.FlushMulus(fileID); err != nil {
		return fmt.Errorf("fail flush mulus for file %d: %w", fileID, err)
	}

	return nil
}

func (s *Storage) deleteData(ctx context.Context, fileName, musAddress string) error {
	resp, err := s.castra[musAddress].client.DeleteData(
		ctx,
		&api.DeleteRequest{
			Name: fileName,
		},
	)

	if err != nil {
		return err
	}

	if resp.Name != fileName {
		return fmt.Errorf("wrong part name %s", resp.Name)
	}

	return nil
}
