CREATE TABLE [dbo].[profileStore] (
    [_id]                 INT            IDENTITY (1, 1) NOT NULL,
    [profileID]           NVARCHAR (MAX) NOT NULL,
    [code]                NVARCHAR (MAX) NOT NULL,
    [name]                NVARCHAR (MAX) NOT NULL,
    [startDate]           NVARCHAR (MAX) NOT NULL,
    [endDate]             NVARCHAR (MAX) NULL,
    [DefaultReleases]     NVARCHAR (MAX) NOT NULL,
    [DefaultReleaseHours] NVARCHAR (MAX) NOT NULL,
    [BlendedRate]         NVARCHAR (MAX) NOT NULL,
    [Rounding]            NVARCHAR (MAX) NOT NULL,
    [HoursPerDay]         NVARCHAR (MAX) NOT NULL,
    [REQPerc]             NVARCHAR (MAX) NULL,
    [ANAPerc]             NVARCHAR (MAX) NULL,
    [DOCPerc]             NVARCHAR (MAX) NULL,
    [PMPerc]              NVARCHAR (MAX) NULL,
    [UATPerc]             NVARCHAR (MAX) NULL,
    [GTMPerc]             NVARCHAR (MAX) NULL,
    [SupportUplift]       NVARCHAR (MAX) NULL,
    [ContigencyPerc]      NVARCHAR (MAX) NULL,
    [_created]            NVARCHAR (MAX) NULL,
    [_createdBy]          NVARCHAR (MAX) NULL,
    [_createdHost]        NVARCHAR (MAX) NULL,
    [_updated]            NVARCHAR (MAX) NULL,
    [_updatedBy]          NVARCHAR (MAX) NULL,
    [_updatedHost]        NVARCHAR (MAX) NULL,
    [_deleted]            NVARCHAR (MAX) NULL,
    [_deletedBy]          NVARCHAR (MAX) NULL,
    [_deletedHost]        NVARCHAR (MAX) NULL,
    [_activity]           NVARCHAR (MAX) NULL,
    [notes]               NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_profileStore] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'profileStore', @level2type = N'COLUMN', @level2name = N'_id';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'profileStore', @level2type = N'COLUMN', @level2name = N'_createdBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'profileStore', @level2type = N'COLUMN', @level2name = N'_updated';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'profileStore', @level2type = N'COLUMN', @level2name = N'_updatedHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'profileStore', @level2type = N'COLUMN', @level2name = N'_deletedBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'Profile of Estimate/Defaults etc.', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'profileStore';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'profileStore', @level2type = N'COLUMN', @level2name = N'_created';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'profileStore', @level2type = N'COLUMN', @level2name = N'_deletedHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'profileStore', @level2type = N'COLUMN', @level2name = N'_createdHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'profileStore', @level2type = N'COLUMN', @level2name = N'_deleted';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'profileStore', @level2type = N'COLUMN', @level2name = N'_updatedBy';


GO

