
# How to properly write SQL codes in golang
Below code is referred to cockroachdb/examples-go. Original codes you can check here: https://github.com/cockroachdb/examples-go/blob/master/photos/db.go

## Declare error in golang 
  var errNoUser = errors.New("no user found")
  var errNoPhoto = errors.New("no photos found")


## Declare queries as constant variable in golang
notice that there is a comma after each column is declared -- sqlite3 grammar.
And there is semi-colon at the end of query. 

  const (
    // TODO(spencer): update the CREATE DATABASE statement in the schema
    //   to pull out the database specified in the DB URL and use it instead
    //   of "photos" below.

    // Create a table schema with back quote -> ' 
    photosSchema = `
  CREATE DATABASE IF NOT EXISTS photos;

  CREATE TABLE IF NOT EXISTS users (
    id           INT,
    photoCount   INT,
    commentCount INT,
    name         STRING,
    address      STRING,

    PRIMARY KEY (id)
  );

  CREATE TABLE IF NOT EXISTS photos (
    id           BYTES DEFAULT uuid_v4(),
    userID       INT,
    commentCount INT,
    caption      STRING,
    latitude     FLOAT,
    longitude    FLOAT,
    timestamp    TIMESTAMP,

    PRIMARY KEY (id),
    UNIQUE INDEX byUserID (userID, timestamp)
  );

  CREATE TABLE IF NOT EXISTS comments (
    -- length check guards against insertion of empty photo ID.
    -- TODO(bdarnell): consider replacing length check with foreign key.
    -- Start with the length check because it's local; we'll want to keep
    -- an eye on performance when introducing the FK.
    photoID   BYTES CHECK (length(photoID) = 16),
    commentID BYTES DEFAULT uuid_v4(),
    userID    INT,
    message   STRING,
    timestamp TIMESTAMP,

    PRIMARY KEY (photoID, timestamp, commentID)
  );`
  )

## Declare query-related function to manipulate data in golang
Divide functions by each step of database - open, table initialization, drop, 
  // openDB opens the database connection according to the context.
  func openDB(cfg Config) (*sql.DB, error) {
    return sql.Open("postgres", cfg.DBUrl)
  }

  // initSchema creates the database schema if it doesn't exist.
  func initSchema(ctx context.Context, db *sql.DB) error {
    _, err := db.ExecContext(ctx, photosSchema)
    return err
  }

  // dropDatabase drops the database.
  func dropDatabase(ctx context.Context, db *sql.DB) error {
    _, err := db.ExecContext(ctx, "DROP DATABASE IF EXISTS photos;")
    return err
  }

// 