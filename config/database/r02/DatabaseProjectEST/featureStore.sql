CREATE TABLE [dbo].[featureStore] (
    [_id]                     INT            IDENTITY (1, 1) NOT NULL,
    [featureID]               NVARCHAR (MAX) NOT NULL,
    [estimationSessionID]     NVARCHAR (MAX) NOT NULL,
    [confidenceID]            NVARCHAR (MAX) NOT NULL,
    [name]                    NVARCHAR (MAX) NOT NULL,
    [devEstimate]             NVARCHAR (MAX) NOT NULL,
    [devUplift]               NVARCHAR (MAX) NULL,
    [reqs]                    NVARCHAR (MAX) NULL,
    [analystTest]             NVARCHAR (MAX) NULL,
    [docs]                    NVARCHAR (MAX) NULL,
    [mgt]                     NVARCHAR (MAX) NULL,
    [uatSupport]              NVARCHAR (MAX) NULL,
    [marketing]               NVARCHAR (MAX) NULL,
    [contingency]             NVARCHAR (MAX) NULL,
    [trackerID]               NVARCHAR (MAX) NULL,
    [adoID]                   NVARCHAR (MAX) NULL,
    [freshdeskID]             NVARCHAR (MAX) NULL,
    [extRef]                  NVARCHAR (MAX) NULL,
    [extRef2]                 NVARCHAR (MAX) NULL,
    [_created]                NVARCHAR (MAX) NULL,
    [_createdBy]              NVARCHAR (MAX) NULL,
    [_createdHost]            NVARCHAR (MAX) NULL,
    [_updated]                NVARCHAR (MAX) NULL,
    [_updatedBy]              NVARCHAR (MAX) NULL,
    [_updatedHost]            NVARCHAR (MAX) NULL,
    [_deleted]                NVARCHAR (MAX) NULL,
    [_deletedBy]              NVARCHAR (MAX) NULL,
    [_deletedHost]            NVARCHAR (MAX) NULL,
    [developer]               NVARCHAR (MAX) NULL,
    [approver]                NVARCHAR (MAX) NULL,
    [notes]                   NVARCHAR (MAX) NULL,
    [offProfile]              NVARCHAR (MAX) NULL,
    [offProfileJustification] NVARCHAR (MAX) NULL,
    [_activity]               NVARCHAR (MAX) NULL,
    [dfReqs]                  NVARCHAR (MAX) NULL,
    [dfAnalystTest]           NVARCHAR (MAX) NULL,
    [dfDocs]                  NVARCHAR (MAX) NULL,
    [dfmgt]                   NVARCHAR (MAX) NULL,
    [dfuatSupport]            NVARCHAR (MAX) NULL,
    [dfmarketing]             NVARCHAR (MAX) NULL,
    [dfcontingency]           NVARCHAR (MAX) NULL,
    [dfdevUplift]             NVARCHAR (MAX) NULL,
    [total]                   NVARCHAR (MAX) NULL,
    [_dbVersion]              NVARCHAR (MAX) NULL,
    [comments]                NVARCHAR (MAX) NULL,
    [description]             NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_featureStore] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'featureStore', @level2type = N'COLUMN', @level2name = N'_deletedHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'featureStore', @level2type = N'COLUMN', @level2name = N'_createdHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'A specific line item in the project', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'featureStore';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'featureStore', @level2type = N'COLUMN', @level2name = N'_updated';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'featureStore', @level2type = N'COLUMN', @level2name = N'_id';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'featureStore', @level2type = N'COLUMN', @level2name = N'_updatedBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'featureStore', @level2type = N'COLUMN', @level2name = N'_updatedHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'featureStore', @level2type = N'COLUMN', @level2name = N'_created';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'featureStore', @level2type = N'COLUMN', @level2name = N'_deleted';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'featureStore', @level2type = N'COLUMN', @level2name = N'_createdBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'featureStore', @level2type = N'COLUMN', @level2name = N'_deletedBy';


GO

