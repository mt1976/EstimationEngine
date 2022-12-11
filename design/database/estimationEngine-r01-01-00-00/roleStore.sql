CREATE TABLE [dbo].[roleStore] (
    [_id]          INT            IDENTITY (1, 1) NOT NULL,
    [Id]           NVARCHAR (MAX) NOT NULL,
    [Name]         NVARCHAR (MAX) NOT NULL,
    [_createdBy]   NVARCHAR (MAX) NULL,
    [_createdHost] NVARCHAR (MAX) NULL,
    [_updatedBy]   NVARCHAR (MAX) NULL,
    [_updatedHost] NVARCHAR (MAX) NULL,
    [_updated]     NVARCHAR (MAX) NULL,
    [_created]     NVARCHAR (MAX) NULL,
    [_deleted]     NVARCHAR (MAX) NULL,
    [_deletedBy]   NVARCHAR (MAX) NULL,
    [_deletedHost] NVARCHAR (MAX) NULL,
    [_dbVersion]   NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_roleStore] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

