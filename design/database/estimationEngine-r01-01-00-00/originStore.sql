CREATE TABLE [dbo].[originStore] (
    [_id]          INT            IDENTITY (1, 1) NOT NULL,
    [originID]     NVARCHAR (MAX) NOT NULL,
    [stateID]      NVARCHAR (MAX) NOT NULL,
    [docTypeID]    NVARCHAR (MAX) NOT NULL,
    [code]         NVARCHAR (MAX) NOT NULL,
    [fullName]     NVARCHAR (MAX) NOT NULL,
    [rate]         NVARCHAR (MAX) NOT NULL,
    [notes]        NVARCHAR (MAX) NULL,
    [startDate]    NVARCHAR (MAX) NULL,
    [endDate]      NVARCHAR (MAX) NULL,
    [_created]     NVARCHAR (MAX) NULL,
    [_createdBy]   NVARCHAR (MAX) NULL,
    [_createdHost] NVARCHAR (MAX) NULL,
    [_updated]     NVARCHAR (MAX) NULL,
    [_updatedBy]   NVARCHAR (MAX) NULL,
    [_updatedHost] NVARCHAR (MAX) NULL,
    [_deleted]     NVARCHAR (MAX) NULL,
    [_deletedBy]   NVARCHAR (MAX) NULL,
    [_deletedHost] NVARCHAR (MAX) NULL,
    [_activity]    NVARCHAR (MAX) NULL,
    [currency]     NVARCHAR (MAX) NULL,
    [_dbVersion]   NVARCHAR (MAX) NULL,
    [comments]     NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_originStore] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'originStore', @level2type = N'COLUMN', @level2name = N'_updatedHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'originStore', @level2type = N'COLUMN', @level2name = N'_updatedBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'originStore', @level2type = N'COLUMN', @level2name = N'_updated';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'originStore', @level2type = N'COLUMN', @level2name = N'_createdHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'originStore', @level2type = N'COLUMN', @level2name = N'_createdBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'originStore', @level2type = N'COLUMN', @level2name = N'_id';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'Customer Data, State etc,', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'originStore';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'originStore', @level2type = N'COLUMN', @level2name = N'_created';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'originStore', @level2type = N'COLUMN', @level2name = N'_deletedHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'originStore', @level2type = N'COLUMN', @level2name = N'_deletedBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'originStore', @level2type = N'COLUMN', @level2name = N'_deleted';


GO

