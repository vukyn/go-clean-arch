package repository

import (
	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/internal/constants"
	"boilerplate-clean-arch/internal/models"
	"boilerplate-clean-arch/internal/session"
	"boilerplate-clean-arch/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// Session repository
type sessionRepo struct {
	cfg         *config.Config
	redisClient *redis.Client
}

// Session repository constructor
func NewSessionRepository(cfg *config.Config, redisClient *redis.Client) session.SessRepository {
	return &sessionRepo{
		cfg,
		redisClient,
	}
}

// Create session in redis
func (s *sessionRepo) CreateSession(ctx context.Context, prefix string, sess *models.Session, expire int) (string, error) {
	log.SetPrefix("[CreateSession]")

	sess.SessionId = uuid.New().String()
	sessionKey := s.createKey(sess.SessionId, prefix)

	sessBytes, err := json.Marshal(&sess)
	if err != nil {
		log.Errorf("marshal failed: %s", err)
		return "", utils.NewError(constants.ERROR_CODE_INTERNAL_SERVER, constants.ERROR_MESSAGE_JSON_MARSHAL)
	}
	if err = s.redisClient.Set(ctx, sessionKey, sessBytes, (time.Second * time.Duration(expire))).Err(); err != nil {
		log.Errorf("set redis failed: %s", err)
		return "", utils.NewError(constants.ERROR_CODE_INTERNAL_SERVER, constants.ERROR_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return sessionKey, nil
}

// Get session by id
func (s *sessionRepo) GetSessionByID(ctx context.Context, sessionId string) (*models.Session, error) {

	sessBytes, err := s.redisClient.Get(ctx, sessionId).Bytes()
	if err != nil {
		log.Errorf("get redis failed: %s", err)
		return nil, utils.NewError(constants.ERROR_CODE_INTERNAL_SERVER, constants.ERROR_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	sess := &models.Session{}
	if err = json.Unmarshal(sessBytes, &sess); err != nil {
		log.Errorf("unmarshal failed: %s", err)
		return nil, utils.NewError(constants.ERROR_CODE_INTERNAL_SERVER, constants.ERROR_MESSAGE_JSON_UNMARSHAL)
	}
	return sess, nil
}

// Delete session by id
func (s *sessionRepo) DeleteByID(ctx context.Context, sessionId string) error {

	if err := s.redisClient.Del(ctx, sessionId).Err(); err != nil {
		log.Errorf("del redis failed: %s", err)
		return utils.NewError(constants.ERROR_CODE_INTERNAL_SERVER, constants.ERROR_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	return nil
}

func (s *sessionRepo) createKey(sessionId, prefix string) string {
	return fmt.Sprintf("%s: %s", prefix, sessionId)
}
