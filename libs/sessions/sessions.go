package sessions

import (
	"errors"
	"os"
	"time"

	"github.com/fredele20/social-media/database"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/sirupsen/logrus"
)

var (
	ErrInvalidUnitOfValidity = errors.New("invalid unit of validity, provide a MINUTE or HOUR")
)

type SessionManager struct {
	logger       *logrus.Logger
	JwtSecretKey string
	db           database.Datastore
}

func NewSessionManager(logger *logrus.Logger, JwtSecretKey string, db database.Datastore) *SessionManager {
	return &SessionManager{
		logger:       logger,
		db:           db,
		JwtSecretKey: JwtSecretKey,
	}
}

func (sm *SessionManager) generatedToken(id, email string) string {
	payload := &TokenPayload{
		Payload: jwt.Payload{
			Issuer:   "Victor",
			Subject:  "Social media app API Token",
			Audience: jwt.Audience{"social-media.com"},
			IssuedAt: jwt.NumericDate(time.Now()),
			JWTID:    "Social-media",
		},
		Id:    id,
		Email: email,
	}

	token, err := jwt.Sign(payload, jwt.NewHS256([]byte(os.Getenv("JWT_SECRETKEY"))))
	if err != nil {
		sm.logger.Debugf("error generating JWT Token: %s", err)
		return ""
	}

	return string(token)
}

func (sm *SessionManager) newSession(userId, email string, validity time.Duration, unitOfValidity UnitOfValidity) *Session {
	token := sm.generatedToken(userId, email)
	return &Session{
		Token:          token,
		UserID:         userId,
		Email:          email,
		Validity:       validity,
		UnitOfValidity: unitOfValidity,
		TimeCreated:    time.Now(),
	}
}

func (sm *SessionManager) CreateSession(payload Session) (string, error) {
	if !payload.UnitOfValidity.IsValid() {
		return "", ErrInvalidUnitOfValidity
	}

	session := sm.newSession(payload.UserID, payload.Email, payload.Validity, payload.UnitOfValidity)

	if err := sm.db.SetSession(session); err != nil {
		sm.logger.WithError(err).Error("failed to persist session to mongodb")
		return "", err
	}

	return session.Token, nil
}

func (sm *SessionManager) DestroySession(key string) error {
	_ = sm.db.ClearSession(key)

	return nil
}
