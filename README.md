# kkreport
[![Build Status](https://travis-ci.org/drkaka/kknotis.svg)](https://travis-ci.org/drkaka/kknotis)
[![Coverage Status](https://codecov.io/github/drkaka/kknotis/coverage.svg?branch=master)](https://codecov.io/github/drkaka/kknotis?branch=master) 

The abuse report module for golang project.

## Database
It is using PostgreSQL as the database and will create a table:

```sql  
CREATE TABLE IF NOT EXISTS report (
	id uuid primary key,
	userid integer,
    at integer,
    handle boolean DEFAULT false,
    reason smallint,
    value text
);
```

## Dependence

```Go
go get github.com/jackc/pgx
go get github.com/satori/go.uuid
```

## Usage 

####First need to use the module with the pgx pool passed in:
```Go
err := kkreport.Use(pool)
```

####Insert a report:
```Go
err := kkreport.InsertReport(3, 0, "value");
```

####Get all reports:
```Go
reports, err := kkreport.GetAllReports(0);
```

####Mark one report as handled:
```Go
err := kkreport.HandleReport(id);
```

####Delete a report record:
```Go
err := kkreport.DeleteReport(id);
```

####Get unhandled reports:
```Go
reports, err := kkreport.GetUnhandledReports(0);
```

####Get handled reports:
```Go
reports, err := kkreport.GetHandledReports(0);
```