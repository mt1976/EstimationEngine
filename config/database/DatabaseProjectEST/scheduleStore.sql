CREATE TABLE [dbo].[scheduleStore] (
    [_id]          INT            IDENTITY (1, 1) NOT NULL,
    [id]           NVARCHAR (MAX) NOT NULL,
    [name]         NVARCHAR (MAX) NULL,
    [description]  NVARCHAR (MAX) NULL,
    [schedule]     NVARCHAR (MAX) NULL,
    [started]      NVARCHAR (MAX) NULL,
    [lastrun]      NVARCHAR (MAX) NULL,
    [message]      NVARCHAR (MAX) NULL,
    [_created]     NVARCHAR (MAX) NULL,
    [_updated]     NVARCHAR (MAX) NULL,
    [type]         NVARCHAR (MAX) NULL,
    [_createdBy]   NVARCHAR (MAX) NULL,
    [_createdHost] NVARCHAR (MAX) NULL,
    [_updatedBy]   NVARCHAR (MAX) NULL,
    [_updatedHost] NVARCHAR (MAX) NULL,
    [human]        NVARCHAR (MAX) NULL,
    [deletedBy]    NVARCHAR (MAX) NULL,
    [deleted]      NVARCHAR (MAX) NULL,
    [deletedHost]  NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_scheduleStore] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

