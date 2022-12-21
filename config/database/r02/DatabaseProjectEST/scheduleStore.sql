CREATE TABLE [dbo].[scheduleStore] (
    [_id]          INT            IDENTITY (1, 1) NOT NULL,
    [id]           NVARCHAR (MAX) NOT NULL,
    [name]         NVARCHAR (MAX) NULL,
    [description]  NVARCHAR (MAX) NULL,
    [schedule]     NVARCHAR (MAX) NULL,
    [started]      NVARCHAR (MAX) NULL,
    [lastrun]      NVARCHAR (MAX) NULL,
    [message]      NVARCHAR (MAX) NULL,
    [type]         NVARCHAR (MAX) NULL,
    [human]        NVARCHAR (MAX) NULL,
    [_created]     NVARCHAR (MAX) NULL,
    [_createdBy]   NVARCHAR (MAX) NULL,
    [_createdHost] NVARCHAR (MAX) NULL,
    [_updated]     NVARCHAR (MAX) NULL,
    [_updatedBy]   NVARCHAR (MAX) NULL,
    [_updatedHost] NVARCHAR (MAX) NULL,
    [_deleted]     NVARCHAR (MAX) NULL,
    [_deletedBy]   NVARCHAR (MAX) NULL,
    [_deletedHost] NVARCHAR (MAX) NULL,
    [_dbVersion]   NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_scheduleStore] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

