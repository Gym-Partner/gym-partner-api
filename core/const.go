package core

const (
	ErrDBUserNotFound = "[USER][REPOSITORY] User not found in database"
	ErrDBCreateUser   = "[USER][REPOSITORY] Failed to create user in database"
	ErrDBGetAllUser   = "[USER][REPOSITORY] Failed to recover all users"
	ErrDBGetOneUser   = "[USER][REPOSITORY] Failed to recover user [%s]"
	ErrDBUpdateUser   = "[USER][REPOSITORY] Failed to update user [%s]"
	ErrDBDeleteUser   = "[USER][REPOSITORY] Failed to delete user [%s]"
	ErrDBUserExist    = "[USER][REPOSITORY] User already exist"

	ErrDBCreateWorkout        = "[WORKOUT][REPOSITORY] Failed to create user's workout"
	ErrDBCreateUnityOfWorkout = "[UNITY OF WORKOUT][REPOSITORY] Failed to create workout's unity"
	ErrDBCreateExercice       = "[EXERCICE][REPOSITORY] Failed to create unity's exercice"
	ErrDBCreateSerie          = "[SERIE][REPOSITORY] Failed to create unity's serie"

	ErrDBGetWorkout        = "[WORKOUT][REPOSITORY] Failed to recover user's workout with his id"
	ErrDBGetUnityOfWorkout = "[UNITY OF WORKOUT][REPOSITORY] Failed to recover workout's unity with his id"
	ErrDBGetExercice       = "[EXERCICE][REPOSITORY] Failed to recover unity's exercice with his id"
	ErrDBGetSerie          = "[SERIE][REPOSITORY] Failed to recover unity's serie with his id"

	ErrIntUserExist     = "[USER][INTERACTOR] User already exist in database"
	ErrIntCreateUserAWS = "[USER][INTERACTOR] Failed to create user in AWS Cognito service"
	ErrIntUserNotExist  = "[USER][INTERACTOR] User %s not found, or not exist in database"

	ErrAWSCognitoCreateSession  = "[AWS][COGNITO] Failed to create new session"
	ErrAWSCognitoCreateUser     = "[AWS][COGNITO] Failed to create user"
	ErrAWSCognitoAuthUser       = "[AWS][COGNITO] Failed to authentificate user"
	ErrAWSCognitoGetUserByToken = "[AWS][COGNITO] Failed to recover the user by his token"
	ErrAWSCognitoDeleteUser     = "[AWS][COGNITO] Failed to delete user"

	ErrConnectDatabase = "[DB CONNECT][PostgreSQL] "
	ErrPingDatabase    = "[DB PING][PostgreSQL] "
	ErrMigrateModel    = "[DB][Postgres] %s"

	ErrEnvParseStart   = "[ENV] Error while parsing START argument: "
	ErrEnvLoad         = "[ENV] Error while loading config file: "
	ErrEnvNoStart      = "[ENV] Error no START parameter provided"
	ErrEnvNoConfigFile = "[ENV] Error no config file found at: %s/config.yaml"
)

const (
	InfoPingDatabase = "[DB PING][PostgreSQL] Connected to the database"
)

const (
	TestINTCreateSuccess = "[TEST][INTERACTOR]CREATE_SUCCESS"
	TestINTGetAllSuccess = "[TEST][INTERACTOR]GET_ALL_SUCCESS"
	TestINTGetOneSuccess = "[TEST][INTERACTOR]GET_ONE_SUCCESS"
	TestINTUdateSuccess  = "[TEST][INTERACTOR]UPDATE_SUCCESS"
	TestINTDeleteSuccess = "[TEST][INTERACTOR]DELETE_SUCCESS"
	TestINTLoginSuccess  = "[TEST][INTERACTOR]LOGIN_SUCCESS"

	TestREPCreateSuccess = "[TEST][REPOSITORY]CREATE_SUCCESS"

	TestUserExistFailed       = "[TEST]USER_EXIST_FAILED"
	TestUserNotExistFailed    = "[TEST]USER_NOT_EXIST_FAILED"
	TestInternalErrorFailed   = "[TEST]INTERNAL_ERROR_FAILED"
	TestUsersNotFound         = "[TEST]USERS_NOT_FOUND_FAILED"
	TestUserNotFound          = "[TEST]USER_NOT_FOUND_FAILED"
	TestUserNotDeletedCognito = "[TEST]USER_NOT_DELETED_COGNITO_FAILED"
	TestUserCreateFailed      = "[TEST]USER_CREATE_FAILED"
)
