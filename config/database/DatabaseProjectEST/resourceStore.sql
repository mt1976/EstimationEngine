CREATE TABLE [dbo].[resourceStore] (
    [_id]          INT            IDENTITY (1, 1) NOT NULL,
    [resourceID]   NVARCHAR (MAX) NOT NULL,
    [code]         NVARCHAR (MAX) NOT NULL,
    [name]         NVARCHAR (MAX) NOT NULL,
    [email]        NVARCHAR (MAX) NOT NULL,
    [class]        NVARCHAR (MAX) NOT NULL,
    [_created]     NVARCHAR (MAX) NULL,
    [_createdBy]   NVARCHAR (MAX) NULL,
    [_createdHost] NVARCHAR (MAX) NULL,
    [_updated]     NVARCHAR (MAX) NULL,
    [_updatedBy]   NVARCHAR (MAX) NULL,
    [_updatedHost] NVARCHAR (MAX) NULL,
    [_deleted]     NVARCHAR (MAX) NULL,
    [_deletedBy]   NVARCHAR (MAX) NULL,
    [_deletedHost] NVARCHAR (MAX) NULL,
    [userActive]   NVARCHAR (MAX) NULL,
    [_activity]    NVARCHAR (MAX) NULL,
    [notes]        NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_resourceStore] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

