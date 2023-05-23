package dataservice

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"

	"protected-storage-server/internal/entity"
	"protected-storage-server/internal/repositories/datarepository"
	"protected-storage-server/internal/security"
)

var _ StorageService = &storageServiceImpl{}

type storageServiceImpl struct {
	rawDataRepository datarepository.RawDataRepository
	cipherManager     *security.CipherManager
}

// New конструктор UserService
func New(rawDataRepository datarepository.RawDataRepository, cipherManager *security.CipherManager) StorageService {
	return &storageServiceImpl{
		rawDataRepository,
		cipherManager,
	}
}

// SaveRawData метод для сохранения произвольных текстовых данных
func (s storageServiceImpl) SaveRawData(ctx context.Context, name, data, userID string) error {
	log.Info().Msgf("dataservice: save raw data for user with ID %s", userID)
	return s.encryptAndSaveData(ctx, name, userID, []byte(data), entity.RAW)
}

// GetRawData метод для получения произвольных текстовых данных
func (s storageServiceImpl) GetRawData(ctx context.Context, name, userID string) (string, error) {
	log.Info().Msgf("dataservice: get raw data with name %s for user with ID %s", name, userID)

	decryptData, err := s.getAndDecryptData(ctx, name, userID, entity.RAW)
	if err != nil {
		return "", err
	}

	return string(decryptData), nil
}

// SaveLoginWithPassword метод для сохранения логина и пароля
func (s storageServiceImpl) SaveLoginWithPassword(ctx context.Context, name, login, password, userID string) error {
	log.Info().Msgf("dataservice: save login with password for user with ID %s", userID)
	cred := entity.CredentialsDTO{
		Login:    login,
		Password: password,
	}

	marshalledCred, err := json.Marshal(cred)
	if err != nil {
		return err
	}
	return s.encryptAndSaveData(ctx, name, userID, marshalledCred, entity.CRED)
}

// GetLoginWithPassword метод для получения логина и пароля
func (s storageServiceImpl) GetLoginWithPassword(ctx context.Context, name, userID string) (entity.CredentialsDTO, error) {
	log.Info().Msgf("dataservice: get credentials with name %s for user with ID %s", name, userID)

	decryptData, err := s.getAndDecryptData(ctx, name, userID, entity.CRED)
	if err != nil {
		return entity.CredentialsDTO{}, err
	}

	cred := entity.CredentialsDTO{}
	if err := json.Unmarshal(decryptData, &cred); err != nil {
		return entity.CredentialsDTO{}, err
	}
	return cred, nil
}

// SaveBinaryData метод для сохранения бинарных данных
func (s storageServiceImpl) SaveBinaryData(ctx context.Context, name string, data []byte, userID string) error {
	log.Info().Msgf("dataservice: save binary data for user with ID %s", userID)
	return s.encryptAndSaveData(ctx, name, userID, data, entity.FILE)
}

// GetBinaryData метод для получения логина и пароля
func (s storageServiceImpl) GetBinaryData(ctx context.Context, name, userID string) ([]byte, error) {
	log.Info().Msgf("dataservice: get binary data with name %s for user with ID %s", name, userID)
	return s.getAndDecryptData(ctx, name, userID, entity.FILE)
}

// SaveCardData метод для сохранения данных банковской карты
func (s storageServiceImpl) SaveCardData(ctx context.Context, name string, cardData entity.CardDataDTO, userID string) error {
	log.Info().Msgf("dataservice: save card data for user with ID %s", userID)

	marshalledCardData, err := json.Marshal(cardData)
	if err != nil {
		return err
	}

	return s.encryptAndSaveData(ctx, name, userID, marshalledCardData, entity.CARD)
}

// GetCardData метод для получения данных банковской карты
func (s storageServiceImpl) GetCardData(ctx context.Context, name, userID string) (entity.CardDataDTO, error) {
	log.Info().Msgf("dataservice: get card data with name %s for user with ID %s", name, userID)

	decryptData, err := s.getAndDecryptData(ctx, name, userID, entity.CARD)
	if err != nil {
		return entity.CardDataDTO{}, err
	}

	card := entity.CardDataDTO{}
	if err = json.Unmarshal(decryptData, &card); err != nil {
		return entity.CardDataDTO{}, err
	}

	return card, nil
}

// GetAllSavedDataNames метод для получения всех названий сохранений
func (s storageServiceImpl) GetAllSavedDataNames(ctx context.Context, userID string) ([]string, error) {
	log.Info().Msgf("dataservice: get data names for user with ID %s", userID)
	return s.rawDataRepository.GetAllSavedDataNames(ctx, userID)
}

func (s storageServiceImpl) encryptAndSaveData(
	ctx context.Context,
	name, userID string,
	data []byte,
	dataType entity.DataType) error {

	savedData, err := s.cipherManager.Encrypt(data)
	if err != nil {
		return err
	}
	return s.rawDataRepository.Save(ctx, userID, name, savedData, dataType)
}

func (s storageServiceImpl) getAndDecryptData(
	ctx context.Context,
	name, userID string,
	dataType entity.DataType) ([]byte, error) {

	data, err := s.rawDataRepository.GetByNameAndTypeAndUserID(ctx, userID, name, dataType)
	if err != nil {
		return nil, err
	}

	decryptData, err := s.cipherManager.Decrypt(data)
	if err != nil {
		return nil, err
	}

	return decryptData, nil
}
