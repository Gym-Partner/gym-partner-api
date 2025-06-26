package core

const (
	InfoPingDatabase = "[DB PING][PostgreSQL] Connected to the database"

	ErrAppDBCreateUser      = "Failed to create user [%s] in database."
	ErrAppDBGetAllUser      = "Failed to retrieve all users in database."
	ErrAppDBGetOneUser      = "Failed to retrieve user [%s] in database."
	ErrAppDBSearchUsers     = "Failed to search users in database."
	ErrAppDBGetUserImage    = "Failed to retrieve user's image [%s] in database."
	ErrAppDBCreateUserImage = "Failed to create users's image in database."
	ErrAppDBDeleteUserImage = "Failed to delete users's image in database."
	ErrAppDBUpdateUser      = "Failed to update user [%s] in database."
	ErrAppDBDeleteUser      = "Failed to delete user [%s] in database."

	ErrAppDBGetWorkouts             = "Failed to retrieve all user's workouts in database."
	ErrAppDBUpdateWorkouts          = "Failed to update user's workouts in database."
	ErrAppDBUpdateUnitiesOfWorkouts = "Failed to update user's unities of workouts in database."
	ErrAppDBUpdateExercises         = "Failed to update user's exercises of workouts in database."
	ErrAppDBUpdateSeries            = "Failed to update user's series of workouts in database."

	ErrAppDBDeleteWorkouts        = "Failed to delete user's workouts in database."
	ErrAppDBDeleteUnityOfWorkouts = "Failed to delete user's unity of workouts in database."
	ErrAppDBDeleteExercises       = "Failed to delete user's exercises of workouts in database."
	ErrAppDBDeleteSeries          = "Failed to delete user's series of workouts in database."

	ErrAppDBAddFollower    = "Failed to add follower for user [%s] in database."
	ErrAppDBRemoveFollower = "Failed to remove follower for user [%s] in database."
	ErrAppDBGetFollowers   = "Failed to retrieve all followers from user [%s] in database."
	ErrAppDBGetFollowed    = "Failed to retrieve all followed from user [%s] in database."

	ErrAppINTUserExist           = "User [%s] already exist in database."
	ErrAppINTUserNotExist        = "User [%s] not found, or not exist in database."
	ErrAppINTUserBindData        = "Failed to bind users data to model"
	ErrAppINTUserImageNotFound   = "User [%s] image not available in request body"
	ErrAppINTUserImageNotOpen    = "User [%s] error to open file image"
	ErrAppINTUserImageUpload     = "User [%s] failed to upload file image"
	ErrAppINTUserImageDeleteS3   = "User [%s] failed to delete file image in S3 service"
	ErrAppINTUserImageDeletePsql = "User [%s] failed to delete file image in PSQL service"

	ErrAppINTWorkoutsNotExist = `User's workout "%s" not exist in database.`

	ErrAppINTFollowerExist    = "Follower [%s] already exist for user [%s] in database."
	ErrAppINTFollowerNotExist = "Follower [%s] not exist for user [%s] in database."

	ErrRabbitMQPublishMessage = "Failed to publish message."

	// ######################################################################################
	// 											LOG
	// ######################################################################################

	// USERS

	ErrDBUserNotFound      = "[USER][REPOSITORY] User not found in database"
	ErrDBCreateUser        = "[USER][REPOSITORY] Failed to create user in database | [ORIGINAL-ERROR] : %s"
	ErrDBGetAllUser        = "[USER][REPOSITORY] Failed to recover all users | [ORIGINAL-ERROR] : %s"
	ErrDBGetOneUser        = "[USER][REPOSITORY] Failed to recover user [%s] | [ORIGINAL-ERROR] : %s"
	ErrDBSearchUsers       = "[USER][REPOSITORY] Failed to search users | [ORIGINAL-ERROR] : %s"
	ErrDBGetUserImage      = "[USER][REPOSITORY] Failed to recover user's image [%s] | [ORIGINAL-ERROR] : %s"
	ErrDBCreateUserImage   = "[USER][REPOSITORY] Failed to create user's image | [ORIGINAL-ERROR] : %s"
	ErrDBUserImageNotFound = "[USER][REPOSITORY] User's image not found in database"
	ErrDBDeleteUserImage   = "[USER][REPOSITORY] Failed to delete user's image [%s] | [ORIGINAL-ERROR] : %s"
	ErrDBUpdateUser        = "[USER][REPOSITORY] Failed to update user [%s] | [ORIGINAL-ERROR] : %s"
	ErrDBDeleteUser        = "[USER][REPOSITORY] Failed to delete user [%s] | [ORIGINAL-ERROR] : %s"

	// WORKOUTS

	ErrDBCreateWorkout        = "[WORKOUT][REPOSITORY] Failed to create user's workout"
	ErrDBCreateUnityOfWorkout = "[UNITY_OF_WORKOUT][REPOSITORY] Failed to create workout's unity"
	ErrDBCreateExercise       = "[EXERCISE][REPOSITORY] Failed to create unity's exercise"
	ErrDBCreateSeries         = "[SERIES][REPOSITORY] Failed to create unity's series"

	ErrDBGetWorkout        = "[WORKOUT][REPOSITORY] Failed to recover user's workout with his id"
	ErrDBGetWorkouts       = "[WORKOUT][REPOSITORY] Failed to recover user's workouts with his id [%s] | [ORIGINAL-ERROR] : %s"
	ErrDBGetUnityOfWorkout = "[UNITY_OF_WORKOUT][REPOSITORY] Failed to recover workout's unity with his id"
	ErrDBGetExercise       = "[EXERCISE][REPOSITORY] Failed to recover unity's exercise with his id"
	ErrDBGetSeries         = "[SERIES][REPOSITORY] Failed to recover unity's series with his id"

	ErrDBUpdateWorkout           = "[WORKOUT][REPOSITORY] Failed to update workout with is user_id: %s | [ORIGINAL-ERROR] : %s"
	ErrDBUpdateUnitiesOfWorkouts = "[UNITIES_OF_WORKOUT][REPOSITORY] Failed to update workout's unities | [ORIGINAL-ERROR] : %s"
	ErrDBUpdateExercises         = "[EXERCISES][REPOSITORY] Failed to update workout's exercises | [ORIGINAL-ERROR] : %s"
	ErrDBUpdateSeries            = "[SERIES][REPOSITORY] Failed to update workout's series | [ORIGINAL-ERROR] : %s"

	ErrDBDeleteWorkout        = "[WORKOUT][REPOSITORY] Failed to delete workout | [ORIGINAL-ERROR] : %s"
	ErrDBDeleteUnityOfWorkout = "[WORKOUT][REPOSITORY] Failed to delete workout's unities | [ORIGINAL-ERROR] : %s"
	ErrDBDeleteExercises      = "[WORKOUT][REPOSITORY] Failed to delete workout's exercises | [ORIGINAL-ERROR] : %s"
	ErrDBDeleteSeries         = "[WORKOUT][REPOSITORY] Failed to delete workout's series | [ORIGINAL-ERROR] : %s"

	// AUTH

	ErrDBCreateAuth = "[AUTH][REPOSITORY] Failed to create auth"

	// FOLLOWERS

	ErrDBAddFollower    = "[FOLLOWER][REPOSITORY] Failed to add follower | [ORIGINAL-ERROR] : %s"
	ErrDBRemoveFollower = "[FOLLOWER][REPOSITORY] Failed to remove follower | [ORIGINAL-ERROR] : %s"
	ErrDBGetFollowers   = "[FOLLOWER][REPOSITORY] Failed to retrieve followers from followed | [ORIGINAL-ERROR] : %s"
	ErrDBGetFollowed    = "[FOLLOWER][REPOSITORY] Failed to retrieve followed from follower | [ORIGINAL-ERROR] : %s"

	// OTHER

	ErrIntUserExist    = "[USER][INTERACTOR] User already exist in database"
	ErrIntUserNotExist = "[USER][INTERACTOR] User %s not found, or not exist in database"

	// DATABASE

	ErrConnectDatabase = "[DB_CONNECT][PostgreSQL] "
	ErrPingDatabase    = "[DB_PING][PostgreSQL] "
	ErrMigrateModel    = "[DB][Postgres] %s"

	// ENV

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
