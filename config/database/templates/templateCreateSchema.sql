IF NOT EXISTS ( SELECT  *
                FROM    sys.schemas
                WHERE   name = N'{{!SQL.SCHEMA}}' )
    EXEC('CREATE SCHEMA {{!SQL.SCHEMA}}');
