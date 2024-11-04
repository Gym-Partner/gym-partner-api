package core

const (
	ErrDBUserNotFound = "[USER][REPOSITORY] User not found in database"
	ErrDBCreateUser   = "[USER][REPOSITORY] Failed to create user in database"
	ErrDBGetAllUser   = "[USER][REPOSITORY] Failed to recover all users"
	ErrDBGetOneUser   = "[USER][REPOSITORY] Failed to recover user [%s]"
	ErrDBUpdateUser   = "[USER][REPOSITORY] Failed to update user [%s]"
	ErrDBDeleteUser   = "[USER][REPOSITORY] Failed to delete user [%s]"
	ErrDBUserExist    = "[USER][REPOSITORY] User already exist"

	ErrIntUserExist     = "[USER][INTERACTOR] User %s alredy exist in database"
	ErrIntInitAWS       = "[USER][INTERACTOR] Failed to init AWS service"
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
	TestCreateSuccess = "[TEST]CREATE_SUCCESS"
	TestGetAllSuccess = "[TEST]GET_ALL_SUCCESS"
	TestGetOneSuccess = "[TEST]GET_ONE_SUCCESS"
	TestUdateSuccess  = "[TEST]UPDATE_SUCCESS"
	TestDeleteSuccess = "[TEST]DELETE_SUCCESS"

	TestUserExistFailed     = "[TEST]USER_EXIST_FAILED"
	TestUserNotExistFailed  = "[TEST]USER_NOT_EXIST_FAILED"
	TestInternalErrorFailed = "[TEST]INTERNAL_ERROR_FAILED"
)
