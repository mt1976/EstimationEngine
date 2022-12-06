CREATE TABLE [dbo].[externalMessageStore] (
    [_id]                    INT            IDENTITY (1, 1) NOT NULL,
    [messageID]              NVARCHAR (MAX) NULL,
    [messageFormat]          NVARCHAR (MAX) NULL,
    [messageDeliveredTo]     NVARCHAR (MAX) NULL,
    [messageBody]            NVARCHAR (MAX) NULL,
    [messageFilename]        NVARCHAR (MAX) NULL,
    [messageLife]            NVARCHAR (MAX) NULL,
    [messageDate]            NVARCHAR (MAX) NULL,
    [messageTime]            NVARCHAR (MAX) NULL,
    [messageTimeoutAction]   NVARCHAR (MAX) NULL,
    [messageACKNAK]          NVARCHAR (MAX) NULL,
    [responseID]             NVARCHAR (MAX) NULL,
    [responseFilename]       NVARCHAR (MAX) NULL,
    [responseBody]           NVARCHAR (MAX) NULL,
    [responseDate]           NVARCHAR (MAX) NULL,
    [responseTime]           NVARCHAR (MAX) NULL,
    [responseAdditionalInfo] NVARCHAR (MAX) NULL,
    [_created]               NVARCHAR (MAX) NULL,
    [_createdBy]             NVARCHAR (MAX) NULL,
    [_createdHost]           NVARCHAR (MAX) NULL,
    [_updated]               NVARCHAR (MAX) NULL,
    [_updatedBy]             NVARCHAR (MAX) NULL,
    [_updatedHost]           NVARCHAR (MAX) NULL,
    [messageTimeout]         NVARCHAR (MAX) NULL,
    [messageEmitted]         NVARCHAR (MAX) NULL,
    [responseRecieved]       NVARCHAR (MAX) NULL,
    [messageClass]           NVARCHAR (MAX) NULL,
    [appID]                  NVARCHAR (MAX) NULL,
    [_deleted]               NVARCHAR (MAX) NULL,
    [_deletedBy]             NVARCHAR (MAX) NULL,
    [_deletedHost]           NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_externalMessageStore] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

