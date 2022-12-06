CREATE TABLE [dbo].[sessionStore] (
    [_id]          INT            IDENTITY (1, 1) NOT NULL,
    [Id]           NVARCHAR (MAX) NULL,
    [Apptoken]     NVARCHAR (MAX) NOT NULL,
    [Createdate]   NVARCHAR (MAX) NULL,
    [Createtime]   NVARCHAR (MAX) NULL,
    [Uniqueid]     NVARCHAR (MAX) NULL,
    [Sessiontoken] NVARCHAR (MAX) NULL,
    [Username]     NVARCHAR (MAX) NULL,
    [Password]     NVARCHAR (MAX) NULL,
    [Userip]       NVARCHAR (MAX) NULL,
    [Userhost]     NVARCHAR (MAX) NULL,
    [Appip]        NVARCHAR (MAX) NULL,
    [Apphost]      NVARCHAR (MAX) NULL,
    [Issued]       NVARCHAR (MAX) NULL,
    [Expiry]       NVARCHAR (MAX) NULL,
    [Expiryraw]    NVARCHAR (MAX) NULL,
    [Expires]      DATETIME2 (7)  NULL,
    [Brand]        NVARCHAR (MAX) NULL,
    [SessionRole]  NVARCHAR (MAX) NULL,
    [_created]     NVARCHAR (MAX) NULL,
    [_createdBy]   NVARCHAR (MAX) NULL,
    [_createdHost] NVARCHAR (MAX) NULL,
    [_updated]     NVARCHAR (MAX) NULL,
    [_updatedBy]   NVARCHAR (MAX) NULL,
    [_updatedHost] NVARCHAR (MAX) NULL,
    [_deleted]     NVARCHAR (MAX) NULL,
    [_deletedBy]   NVARCHAR (MAX) NULL,
    [_deletedHost] NVARCHAR (MAX) NULL,
    PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

