CREATE TABLE [dbo].[estimationSessionStore] (
    [_id]                 INT            IDENTITY (1, 1) NOT NULL,
    [estimationSessionID] NVARCHAR (MAX) NOT NULL,
    [projectID]           NVARCHAR (MAX) NOT NULL,
    [estimationStateID]   NVARCHAR (MAX) NOT NULL,
    [notes]               NVARCHAR (MAX) NOT NULL,
    [releases]            NVARCHAR (MAX) NULL,
    [total]               NVARCHAR (MAX) NULL,
    [contingency]         NVARCHAR (MAX) NULL,
    [reqDays]             NVARCHAR (MAX) NULL,
    [regCost]             NVARCHAR (MAX) NULL,
    [impDays]             NVARCHAR (MAX) NULL,
    [impCost]             NVARCHAR (MAX) NULL,
    [uatDays]             NVARCHAR (MAX) NULL,
    [uatCost]             NVARCHAR (MAX) NULL,
    [mgtDays]             NVARCHAR (MAX) NULL,
    [mgtCost]             NVARCHAR (MAX) NULL,
    [relDays]             NVARCHAR (MAX) NULL,
    [relCost]             NVARCHAR (MAX) NULL,
    [supportUplift]       NVARCHAR (MAX) NULL,
    [_created]            NVARCHAR (MAX) NULL,
    [_createdBy]          NVARCHAR (MAX) NULL,
    [_createdHost]        NVARCHAR (MAX) NULL,
    [_updated]            NVARCHAR (MAX) NULL,
    [_updatedBy]          NVARCHAR (MAX) NULL,
    [_updatedHost]        NVARCHAR (MAX) NULL,
    [_deleted]            NVARCHAR (MAX) NULL,
    [_deletedBy]          NVARCHAR (MAX) NULL,
    [_deletedHost]        NVARCHAR (MAX) NULL,
    [name]                NVARCHAR (MAX) NULL,
    [adoID]               NVARCHAR (MAX) NULL,
    [freshdeskID]         NVARCHAR (MAX) NULL,
    [trackerID]           NVARCHAR (MAX) NULL,
    [estRef]              NVARCHAR (MAX) NULL,
    [extRef]              NVARCHAR (MAX) NULL,
    [_activity]           NVARCHAR (MAX) NULL,
    [_dbVersion]          NVARCHAR (MAX) NULL,
    [comments]            NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_estimationSessionStore] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'Results of an Estimation Session', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'estimationSessionStore';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'estimationSessionStore', @level2type = N'COLUMN', @level2name = N'_created';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'estimationSessionStore', @level2type = N'COLUMN', @level2name = N'_createdHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'estimationSessionStore', @level2type = N'COLUMN', @level2name = N'_deletedHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'estimationSessionStore', @level2type = N'COLUMN', @level2name = N'_updatedBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'estimationSessionStore', @level2type = N'COLUMN', @level2name = N'_deleted';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'estimationSessionStore', @level2type = N'COLUMN', @level2name = N'_id';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'estimationSessionStore', @level2type = N'COLUMN', @level2name = N'_deletedBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'estimationSessionStore', @level2type = N'COLUMN', @level2name = N'_createdBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'estimationSessionStore', @level2type = N'COLUMN', @level2name = N'_updatedHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'estimationSessionStore', @level2type = N'COLUMN', @level2name = N'_updated';


GO

