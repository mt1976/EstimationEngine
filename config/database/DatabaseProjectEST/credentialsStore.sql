CREATE TABLE [dbo].[credentialsStore] (
    [_id]          INT            IDENTITY (1, 1) NOT NULL,
    [Id]           NVARCHAR (MAX) NOT NULL,
    [Username]     NVARCHAR (MAX) NOT NULL,
    [Password]     NVARCHAR (MAX) NULL,
    [Firstname]    NVARCHAR (MAX) NULL,
    [Lastname]     NVARCHAR (MAX) NULL,
    [Knownas]      NVARCHAR (MAX) NULL,
    [Email]        NVARCHAR (MAX) NULL,
    [Issued]       NVARCHAR (MAX) NULL,
    [Expiry]       NVARCHAR (MAX) NULL,
    [RoleType]     NVARCHAR (MAX) NULL,
    [Brand]        NVARCHAR (MAX) NULL,
    [_created]     NVARCHAR (MAX) NULL,
    [_updated]     NVARCHAR (MAX) NULL,
    [_createdBy]   NVARCHAR (MAX) NULL,
    [_createdHost] NVARCHAR (MAX) NULL,
    [_updatedBy]   NVARCHAR (MAX) NULL,
    [_updatedHost] NVARCHAR (MAX) NULL,
    [State]        NVARCHAR (MAX) NULL,
    [Notes]        NVARCHAR (255) NULL,
    [_deleted]     NVARCHAR (MAX) NULL,
    [_deletedBy]   NVARCHAR (MAX) NULL,
    [_deletedHost] NVARCHAR (MAX) NULL,
    [_activity]    NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_credentialsStore] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

