package core

const (
	InfoPingDatabase = "[DB PING][PostgreSQL] Connected to the database"

	ErrDBUserNotFound = "[USER][REPOSITORY] User not found in database"
	ErrDBCreateUser   = "[USER][REPOSITORY] Failed to create user in database"
	ErrDBGetAllUser   = "[USER][REPOSITORY] Failed to recover all users"
	ErrDBGetOneUser   = "[USER][REPOSITORY] Failed to recover user [%s]"
	ErrDBUpdateUser   = "[USER][REPOSITORY] Failed to update user [%s]"
	ErrDBDeleteUser   = "[USER][REPOSITORY] Failed to delete user [%s]"
	ErrDBUserExist    = "[USER][REPOSITORY] User already exist"

	ErrDBCreateWorkout        = "[WORKOUT][REPOSITORY] Failed to create user's workout"
	ErrDBCreateUnityOfWorkout = "[UNITY_OF_WORKOUT][REPOSITORY] Failed to create workout's unity"
	ErrDBCreateExercice       = "[EXERCICE][REPOSITORY] Failed to create unity's exercice"
	ErrDBCreateSerie          = "[SERIE][REPOSITORY] Failed to create unity's serie"

	ErrDBGetWorkout        = "[WORKOUT][REPOSITORY] Failed to recover user's workout with his id"
	ErrDBGetUnityOfWorkout = "[UNITY_OF_WORKOUT][REPOSITORY] Failed to recover workout's unity with his id"
	ErrDBGetExercice       = "[EXERCICE][REPOSITORY] Failed to recover unity's exercice with his id"
	ErrDBGetSerie          = "[SERIE][REPOSITORY] Failed to recover unity's serie with his id"

	ErrDBCreateAuth = "[AUTH][REPOSITORY] Failed to create auth"

	ErrIntUserExist    = "[USER][INTERACTOR] User already exist in database"
	ErrIntUserNotExist = "[USER][INTERACTOR] User %s not found, or not exist in database"

	ErrConnectDatabase = "[DB_CONNECT][PostgreSQL] "
	ErrPingDatabase    = "[DB_PING][PostgreSQL] "
	ErrMigrateModel    = "[DB][Postgres] %s"

	ErrEnvParseStart   = "[ENV] Error while parsing START argument: "
	ErrEnvLoad         = "[ENV] Error while loading config file: "
	ErrEnvNoStart      = "[ENV] Error no START parameter provided"
	ErrEnvNoConfigFile = "[ENV] Error no config file found at: %s/config.yaml"

	// ######################################################################################
	// 										TEST LOG
	// ######################################################################################

	TestINTUserCreateSuccess        = "[TEST][INTERACTOR][USER][SUCCESS]CREATE"
	TestINTUserGetAllSuccess        = "[TEST][INTERACTOR][USER][SUCCESS]GET_ALL"
	TestINTUserGetOneByIdSuccess    = "[TEST][INTERACTOR][USER][SUCCESS]GET_ONE_BY_ID"
	TestINTUserGetOneByEmailSuccess = "[TEST][INTERACTOR][USER][SUCCESS]GET_ONE_BY_EMAIL"
	TestINTUserUdateSuccess         = "[TEST][INTERACTOR][USER][SUCCESS]UPDATE"
	TestINTUserDeleteSuccess        = "[TEST][INTERACTOR][USER][SUCCESS]DELETE"
	TestINTUserLoginSuccess         = "[TEST][INTERACTOR][USER][SUCCESS]LOGIN"

	TestREPUserCreateSuccess         = "[TEST][REPOSITORY][USER][SUCCESS]CREATE"
	TestREPUserGetAllSuccess         = "[TEST[REPOSITORY]][USER][SUCCESS]GET_ALL"
	TestREPUserGetOneByIdSuccess     = "[TEST][REPOSITORY][USER][SUCCESS]GET_ONE_BY_ID"
	TestREPUserGetOneByEmailSuccess  = "[TEST][REPOSITORY][USER][SUCCESS]GET_ONE_BY_EMAIL"
	TestREPUserUpdateSuccess         = "[TEST][REPOSITORY][USER][SUCCESS]UPDATE"
	TestREPUserDeleteSuccess         = "[TEST][REPOSITORY][USER][SUCCESS]DELETE"
	TestREPUserIsExistByIdSuccess    = "[TEST][REPOSITORY][USER][SUCCESS]IS_EXIST_BY_ID"
	TestREPUserIsExistByEmailSuccess = "[TEST][REPOSITORY][USER][SUCCESS]IS_EXIST_BY_EMAIL"

	TestCONUserCreateSuccess = "[TEST][CONTROLLER][USER][SUCCESS]CREATE"
	TestCONUserGetAllSuccess = "[TEST][CONTROLLER][USER][SUCCESS]GET_ALL"
	TestCONUserGetOneSuccess = "[TEST][CONTROLLER][USER][SUCCESS]GET_ONE"
	TestCONUserUpdateSuccess = "[TEST][CONTROLLER][USER][SUCCESS]UPDATE"
	TestCONUserDeleteSuccess = "[TEST][CONTROLLER][USER][SUCCESS]DELETE"
	TestCONUserLoginSuccess  = "[TEST][CONTROLLER][USER][SUCCESS]LOGIN"

	TestUserExistFailed     = "[TEST][FAILED]USER_EXIST"
	TestUserNotExistFailed  = "[TEST][FAILED]USER_NOT_EXIST"
	TestInternalErrorFailed = "[TEST][FAILED]INTERNAL_ERROR"
	TestUsersNotFound       = "[TEST][FAILED]USERS_NOT_FOUND"
	TestUserNotFound        = "[TEST][FAILED]USER_NOT_FOUND"
	TestUserCreateFailed    = "[TEST][FAILED]USER_CREATE"
	TestUserUpdateFailed    = "[TEST][FAILED]USER_UPDATE"
	TestUserDeleteFailed    = "[TEST][FAILED]USER_DELETE"
	TestUserLoginFailed     = "[TEST][FAILED]USER_LOGIN"

	TestINTWorkoutCreateSuccess = "[TEST][INTERACTOR][WORKOUT][SUCCESS]CREATE"
	TestINTWorkoutGetSuccess    = "[TEST][INTERACTOR][WORKOUT][SUCCESS]GET]"

	TestREPWorkoutCreateSuccess        = "[TEST][REPOSITORY][WORKOUT][SUCCESS]CREATE"
	TestREPUnityOfWorkoutCreateSuccess = "[TEST][REPOSITORY][UNITY_OF_WORKOUT][SUCCESS]CREATE"
	TestREPExerciceCreateSuccess       = "[TEST][REPOSITORY][EXERCICE][SUCCESS]CREATE"
	TestREPSerieCreateSuccess          = "[TEST][REPOSITORY][SERIE][SUCCESS]CREATE"
	TestREPWorkoutGetSuccess           = "[TEST][REPOSITORY][WORKOUT][SUCCESS]GET"
	TestREPUnityOfWorkoutGetSuccess    = "[TEST][REPOSITORY][UNITY_OF_WORKOUT][SUCCESS]GET"
	TestREPExerciceGetSuccess          = "[TEST][REPOSITORY][EXERCICE][SUCCESS]GET"
	TESTREPSerieGetSuccess             = "[TEST][REPOSITORY][SERIE][SUCCESS]GET"

	TestCONWorkoutCreateSuccess = "[TEST][CONTROLLER][WORKOUT][SUCCESS]CREATE"
	TestCONWorkoutGetSuccess    = "[TEST][CONTROLLER][WORKOUT][SUCCESS]GET"

	TestWorkoutCreateFailed          = "[TEST][FAILED]WORKOUT_CREATE"
	TestUnitiesOfWorkoutCreateFailed = "[TEST][FAILED]UNITIES_OF_WORKOUT_CREATE"
	TestExercicesCreateFailed        = "[TEST][FAILED]EXCERCICES_CREATE"
	TestSeriesCreateFailed           = "[TEST][FAILED]SERIES_CREATE"
	TestWorkoutGetFailed             = "[TEST][FAILED]WORKOUT_GET"
	TestUnitiesOfWorkoutGetFailed    = "[TEST][FAILED]UNITIES_OF_WORKOUT_GET"
	TestExercicesGetFailed           = "[TEST][FAILED]EXERCICES_GET"
	TestSeriesGetFailed              = "[TEST][FAILED]SERIES_GET"
)
