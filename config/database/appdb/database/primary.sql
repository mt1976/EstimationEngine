IF NOT EXISTS (
        SELECT *
        FROM sys.databases
        WHERE name = '{{!DATABASENAME}}-{{!INSTANCE}}'
        )
BEGIN
    CREATE DATABASE [{{!DATABASENAME}}-{{!INSTANCE}}]
END
