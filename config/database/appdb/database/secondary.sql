IF NOT EXISTS (
        SELECT *
        FROM sys.databases
        WHERE name = '{{!DATABASENAME}}'
        )
BEGIN
    CREATE DATABASE [{{!DATABASENAME}}]
END
