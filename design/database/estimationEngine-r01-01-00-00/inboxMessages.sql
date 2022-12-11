CREATE TABLE [dbo].[inboxMessages] (
    [_id]              INT            IDENTITY (1, 1) NOT NULL,
    [_created]         NVARCHAR (MAX) NULL,
    [_updated]         NVARCHAR (MAX) NULL,
    [_createdBy]       NVARCHAR (MAX) NULL,
    [_createdHost]     NVARCHAR (MAX) NULL,
    [_updatedBy]       NVARCHAR (MAX) NULL,
    [_updatedHost]     NVARCHAR (MAX) NULL,
    [MailId]           NVARCHAR (MAX) NULL,
    [MailTo]           NVARCHAR (MAX) NULL,
    [MailFrom]         NVARCHAR (MAX) NULL,
    [MailSource]       NVARCHAR (MAX) NULL,
    [MailSent]         NVARCHAR (MAX) NULL,
    [MailUnread]       NVARCHAR (MAX) NULL,
    [MailSubject]      NVARCHAR (MAX) NULL,
    [MailContent]      NVARCHAR (MAX) NULL,
    [MailAcknowledged] NVARCHAR (MAX) NULL,
    [_deleted]         NVARCHAR (MAX) NULL,
    [_deletedBy]       NVARCHAR (MAX) NULL,
    [_deletedHost]     NVARCHAR (MAX) NULL,
    [_dbVersion]       NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_inboxMessages] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

