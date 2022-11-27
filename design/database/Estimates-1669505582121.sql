
        
CREATE TABLE confidenceStore
(
  _id          int           NOT NULL IDENTITY(1,1),
  ConfidenceID nvarchar(255) NOT NULL,
  Code         nvarchar(4)   NOT NULL,
  Name         nvarchar(255) NOT NULL,
  Perc         nvarchar(255) NOT NULL,
  _created     nvarchar(255),
  _createdBy   nvarchar(255),
  _createdHost nvarchar(255),
  _updated     nvarchar(255),
  _updatedBy   nvarchar(255),
  _updatedHost nvarchar(255),
  _deleted     nvarchar(255),
  _deletedBy   nvarchar(255),
  _deletedHost nvarchar(255),
  CONSTRAINT PK_confidenceStore PRIMARY KEY (_id)
)
GO

ALTER TABLE confidenceStore
  ADD CONSTRAINT UQ__id UNIQUE (_id)
GO

ALTER TABLE confidenceStore
  ADD CONSTRAINT UQ_ConfidenceID UNIQUE (ConfidenceID)
GO

ALTER TABLE confidenceStore
  ADD CONSTRAINT UQ_Code UNIQUE (Code)
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'Confidence Levels of Development Estimate', 'user', dbo, 'table', 'confidenceStore'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'confidenceStore', 'column', '_id'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'confidenceStore', 'column', '_created'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'confidenceStore', 'column', '_createdBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'confidenceStore', 'column', '_createdHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'confidenceStore', 'column', '_updated'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'confidenceStore', 'column', '_updatedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'confidenceStore', 'column', '_updatedHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'confidenceStore', 'column', '_deleted'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'confidenceStore', 'column', '_deletedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'confidenceStore', 'column', '_deletedHost'
GO

CREATE TABLE docTypeStore
(
  _id          int           NOT NULL IDENTITY(1,1),
  DocTypeID    nvarchar(255) NOT NULL,
  Code         nvarchar(4)   NOT NULL,
  Name         nvarchar(255) NOT NULL,
  _created     nvarchar(255),
  _createdBy   nvarchar(255),
  _createdHost nvarchar(255),
  _updated     nvarchar(255),
  _updatedBy   nvarchar(255),
  _updatedHost nvarchar(255),
  _deleted     nvarchar(255),
  _deletedBy   nvarchar(255),
  _deletedHost nvarchar(255),
  CONSTRAINT PK_docTypeStore PRIMARY KEY (_id)
)
GO

ALTER TABLE docTypeStore
  ADD CONSTRAINT UQ__id UNIQUE (_id)
GO

ALTER TABLE docTypeStore
  ADD CONSTRAINT UQ_DocTypeID UNIQUE (DocTypeID)
GO

ALTER TABLE docTypeStore
  ADD CONSTRAINT UQ_Code UNIQUE (Code)
GO

ALTER TABLE docTypeStore
  ADD CONSTRAINT UQ_Name UNIQUE (Name)
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'Document Types "RSC/NC"', 'user', dbo, 'table', 'docTypeStore'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'docTypeStore', 'column', '_id'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'docTypeStore', 'column', '_created'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'docTypeStore', 'column', '_createdBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'docTypeStore', 'column', '_createdHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'docTypeStore', 'column', '_updated'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'docTypeStore', 'column', '_updatedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'docTypeStore', 'column', '_updatedHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'docTypeStore', 'column', '_deleted'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'docTypeStore', 'column', '_deletedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'docTypeStore', 'column', '_deletedHost'
GO

CREATE TABLE estimationSessionStore
(
  _id                 int           NOT NULL IDENTITY(1,1),
  EstimationSessionID nvarchar(255) NOT NULL,
  SessionStateID      nvarchar(255) NOT NULL,
  Notes               nvarchar(255) NOT NULL,
  Releases            nvarchar(255),
  Total               nvarchar(255),
  Contingency         nvarchar(255),
  REQDays             nvarchar(255),
  REQCost             nvarchar(255),
  IMPDays             nvarchar(255),
  IMPCost             nvarchar(255),
  UATDays             nvarchar(255),
  UATCost             nvarchar(255),
  PMDays              nvarchar(255),
  PMCost              nvarchar(255),
  RELDays             nvarchar(255),
  RELCost             nvarchar(255),
  SupportUplift       nvarchar(255),
  _created            nvarchar(255),
  _createdBy          nvarchar(255),
  _createdHost        nvarchar(255),
  _updated            nvarchar(255),
  _updatedBy          nvarchar(255),
  _updatedHost        nvarchar(255),
  _deleted            nvarchar(255),
  _deletedBy          nvarchar(255),
  _deletedHost        nvarchar(255),
  CONSTRAINT PK_estimationSessionStore PRIMARY KEY (_id)
)
GO

ALTER TABLE estimationSessionStore
  ADD CONSTRAINT UQ__id UNIQUE (_id)
GO

ALTER TABLE estimationSessionStore
  ADD CONSTRAINT UQ_EstimationSessionID UNIQUE (EstimationSessionID)
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'Results of an Estimation Session', 'user', dbo, 'table', 'estimationSessionStore'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'estimationSessionStore', 'column', '_id'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'estimationSessionStore', 'column', '_created'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'estimationSessionStore', 'column', '_createdBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'estimationSessionStore', 'column', '_createdHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'estimationSessionStore', 'column', '_updated'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'estimationSessionStore', 'column', '_updatedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'estimationSessionStore', 'column', '_updatedHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'estimationSessionStore', 'column', '_deleted'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'estimationSessionStore', 'column', '_deletedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'estimationSessionStore', 'column', '_deletedHost'
GO

CREATE TABLE featureStore
(
  _id                 int           NOT NULL IDENTITY(1,1),
  FeatureID           nvarchar(255) NOT NULL,
  EstimationSessionID nvarchar(255) NOT NULL,
  ConfidenceID        nvarchar(255) NOT NULL,
  Name                nvarchar(255) NOT NULL,
  DevEstimate         nvarchar(255) NOT NULL,
  DevUplift           nvarchar(255),
  Reqs                nvarchar(255),
  AnalystTest         nvarchar(255),
  Docs                nvarchar(255),
  Mgt                 nvarchar(255),
  UATSupport          nvarchar(255),
  Marketing           nvarchar(255),
  Contingency         nvarchar(255),
  TrackerID           nvarchar(255),
  ADOID               nvarchar(255),
  FDID                nvarchar(255),
  ExtRef              nvarchar(255),
  _created            nvarchar(255),
  _createdBy          nvarchar(255),
  _createdHost        nvarchar(255),
  _updated            nvarchar(255),
  _updatedBy          nvarchar(255),
  _updatedHost        nvarchar(255),
  _deleted            nvarchar(255),
  _deletedBy          nvarchar(255),
  _deletedHost        nvarchar(255),
  CONSTRAINT PK_featureStore PRIMARY KEY (_id)
)
GO

ALTER TABLE featureStore
  ADD CONSTRAINT UQ__id UNIQUE (_id)
GO

ALTER TABLE featureStore
  ADD CONSTRAINT UQ_FeatureID UNIQUE (FeatureID)
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'A specific line item in the project', 'user', dbo, 'table', 'featureStore'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'featureStore', 'column', '_id'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'featureStore', 'column', '_created'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'featureStore', 'column', '_createdBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'featureStore', 'column', '_createdHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'featureStore', 'column', '_updated'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'featureStore', 'column', '_updatedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'featureStore', 'column', '_updatedHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'featureStore', 'column', '_deleted'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'featureStore', 'column', '_deletedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'featureStore', 'column', '_deletedHost'
GO

CREATE TABLE originStateStore
(
  _id           int           NOT NULL IDENTITY(1,1),
  OriginStateID nvarchar(255) NOT NULL,
  Code          nvarchar(4)   NOT NULL,
  Name          nvarchar(255) NOT NULL,
  _created      nvarchar(255),
  _createdBy    nvarchar(255),
  _createdHost  nvarchar(255),
  _updated      nvarchar(255),
  _updatedBy    nvarchar(255),
  _updatedHost  nvarchar(255),
  _deleted      nvarchar(255),
  _deletedBy    nvarchar(255),
  _deletedHost  nvarchar(255),
  CONSTRAINT PK_originStateStore PRIMARY KEY (_id)
)
GO

ALTER TABLE originStateStore
  ADD CONSTRAINT UQ__id UNIQUE (_id)
GO

ALTER TABLE originStateStore
  ADD CONSTRAINT UQ_OriginStateID UNIQUE (OriginStateID)
GO

ALTER TABLE originStateStore
  ADD CONSTRAINT UQ_Code UNIQUE (Code)
GO

ALTER TABLE originStateStore
  ADD CONSTRAINT UQ_Name UNIQUE (Name)
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'Customer State (Active, Prospect, Past)', 'user', dbo, 'table', 'originStateStore'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStateStore', 'column', '_id'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStateStore', 'column', '_created'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStateStore', 'column', '_createdBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStateStore', 'column', '_createdHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStateStore', 'column', '_updated'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStateStore', 'column', '_updatedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStateStore', 'column', '_updatedHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStateStore', 'column', '_deleted'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStateStore', 'column', '_deletedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStateStore', 'column', '_deletedHost'
GO

CREATE TABLE originStore
(
  _id          int           NOT NULL IDENTITY(1,1),
  OriginID     nvarchar(255) NOT NULL,
  StateID      nvarchar(255) NOT NULL,
  DocTypeID    nvarchar(255) NOT NULL,
  Code         nvarchar(4)   NOT NULL,
  FullName     nvarchar(255) NOT NULL,
  Rate         nvarchar(255) NOT NULL,
  Notes        nvarchar(255),
  StartDate    nvarchar(255),
  EndDate      nvarchar(255),
  _created     nvarchar(255),
  _createdBy   nvarchar(255),
  _createdHost nvarchar(255),
  _updated     nvarchar(255),
  _updatedBy   nvarchar(255),
  _updatedHost nvarchar(255),
  _deleted     nvarchar(255),
  _deletedBy   nvarchar(255),
  _deletedHost nvarchar(255),
  CONSTRAINT PK_originStore PRIMARY KEY (_id)
)
GO

ALTER TABLE originStore
  ADD CONSTRAINT UQ__id UNIQUE (_id)
GO

ALTER TABLE originStore
  ADD CONSTRAINT UQ_OriginID UNIQUE (OriginID)
GO

ALTER TABLE originStore
  ADD CONSTRAINT UQ_Code UNIQUE (Code)
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'Customer Data, State etc,', 'user', dbo, 'table', 'originStore'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStore', 'column', '_id'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStore', 'column', '_created'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStore', 'column', '_createdBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStore', 'column', '_createdHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStore', 'column', '_updated'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStore', 'column', '_updatedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStore', 'column', '_updatedHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStore', 'column', '_deleted'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStore', 'column', '_deletedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'originStore', 'column', '_deletedHost'
GO

CREATE TABLE profileStore
(
  _id                 int           NOT NULL IDENTITY(1,1),
  ProfileID           nvarchar(255) NOT NULL,
  Code                nvarchar(4)   NOT NULL,
  Name                nvarchar(255) NOT NULL,
  StartDate           nvarchar(255) NOT NULL,
  EndDate             nvarchar(255),
  DefaultReleases     nvarchar(255) NOT NULL,
  DefaultReleaseHours nvarchar(255) NOT NULL,
  BlendedRate         nvarchar(255) NOT NULL,
  Rounding            nvarchar(255) NOT NULL,
  HoursPerDay         nvarchar(255) NOT NULL,
  REQPerc             nvarchar(255),
  ANAPerc             nvarchar(255),
  DOCPerc             nvarchar(255),
  PMPerc              nvarchar(255),
  UATPerc             nvarchar(255),
  GTMPerc             nvarchar(255),
  SupportUplift       nvarchar(255),
  ContigencyPerc      nvarchar(255),
  _created            nvarchar(255),
  _createdBy          nvarchar(255),
  _createdHost        nvarchar(255),
  _updated            nvarchar(255),
  _updatedBy          nvarchar(255),
  _updatedHost        nvarchar(255),
  _deleted            nvarchar(255),
  _deletedBy          nvarchar(255),
  _deletedHost        nvarchar(255),
  CONSTRAINT PK_profileStore PRIMARY KEY (_id)
)
GO

ALTER TABLE profileStore
  ADD CONSTRAINT UQ__id UNIQUE (_id)
GO

ALTER TABLE profileStore
  ADD CONSTRAINT UQ_ProfileID UNIQUE (ProfileID)
GO

ALTER TABLE profileStore
  ADD CONSTRAINT UQ_Code UNIQUE (Code)
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'Profile of Estimate/Defaults etc.', 'user', dbo, 'table', 'profileStore'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'profileStore', 'column', '_id'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'profileStore', 'column', '_created'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'profileStore', 'column', '_createdBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'profileStore', 'column', '_createdHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'profileStore', 'column', '_updated'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'profileStore', 'column', '_updatedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'profileStore', 'column', '_updatedHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'profileStore', 'column', '_deleted'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'profileStore', 'column', '_deletedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'profileStore', 'column', '_deletedHost'
GO

CREATE TABLE projectStateStore
(
  _id            int           NOT NULL IDENTITY(1,1),
  ProjectStateID nvarchar(255) NOT NULL,
  Code           nvarchar(4)   NOT NULL,
  Name           nvarchar(255) NOT NULL,
  _created       nvarchar(255),
  _createdBy     nvarchar(255),
  _createdHost   nvarchar(255),
  _updated       nvarchar(255),
  _updatedBy     nvarchar(255),
  _updatedHost   nvarchar(255),
  _deleted       nvarchar(255),
  _deletedBy     nvarchar(255),
  _deletedHost   nvarchar(255),
  CONSTRAINT PK_projectStateStore PRIMARY KEY (_id)
)
GO

ALTER TABLE projectStateStore
  ADD CONSTRAINT UQ__id UNIQUE (_id)
GO

ALTER TABLE projectStateStore
  ADD CONSTRAINT UQ_ProjectStateID UNIQUE (ProjectStateID)
GO

ALTER TABLE projectStateStore
  ADD CONSTRAINT UQ_Code UNIQUE (Code)
GO

ALTER TABLE projectStateStore
  ADD CONSTRAINT UQ_Name UNIQUE (Name)
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'Project Lifecycle', 'user', dbo, 'table', 'projectStateStore'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStateStore', 'column', '_id'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStateStore', 'column', '_created'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStateStore', 'column', '_createdBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStateStore', 'column', '_createdHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStateStore', 'column', '_updated'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStateStore', 'column', '_updatedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStateStore', 'column', '_updatedHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStateStore', 'column', '_deleted'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStateStore', 'column', '_deletedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStateStore', 'column', '_deletedHost'
GO

CREATE TABLE projectStore
(
  _id            int           NOT NULL IDENTITY(1,1),
  WorkID         nvarchar(255) NOT NULL,
  OriginID       nvarchar(255) NOT NULL,
  ProjectStateID nvarchar(255) NOT NULL,
  ProfileID      nvarchar(255) NOT NULL,
  Name           nvarchar(255) NOT NULL,
  Description    nvarchar(255),
  StartDate      nvarchar(255),
  EndDate        nvarchar(255),
  _created       nvarchar(255),
  _createdBy     nvarchar(255),
  _createdHost   nvarchar(255),
  _updated       nvarchar(255),
  _updatedBy     nvarchar(255),
  _updatedHost   nvarchar(255),
  _deleted       nvarchar(255),
  _deletedBy     nvarchar(255),
  _deletedHost   nvarchar(255),
  CONSTRAINT PK_projectStore PRIMARY KEY (_id)
)
GO

ALTER TABLE projectStore
  ADD CONSTRAINT UQ__id UNIQUE (_id)
GO

ALTER TABLE projectStore
  ADD CONSTRAINT UQ_WorkID UNIQUE (WorkID)
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'A customer/internal project', 'user', dbo, 'table', 'projectStore'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStore', 'column', '_id'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStore', 'column', '_created'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStore', 'column', '_createdBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStore', 'column', '_createdHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStore', 'column', '_updated'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStore', 'column', '_updatedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStore', 'column', '_updatedHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStore', 'column', '_deleted'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStore', 'column', '_deletedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'projectStore', 'column', '_deletedHost'
GO

CREATE TABLE sessionStateStore
(
  _id            int           NOT NULL IDENTITY(1,1),
  SessionStateID nvarchar(255) NOT NULL,
  Code           nvarchar(4)   NOT NULL,
  Name           nvarchar(255) NOT NULL,
  _created       nvarchar(255),
  _createdBy     nvarchar(255),
  _createdHost   nvarchar(255),
  _updated       nvarchar(255),
  _updatedBy     nvarchar(255),
  _updatedHost   nvarchar(255),
  _deleted       nvarchar(255),
  _deletedBy     nvarchar(255),
  _deletedHost   nvarchar(255),
  CONSTRAINT PK_sessionStateStore PRIMARY KEY (_id)
)
GO

ALTER TABLE sessionStateStore
  ADD CONSTRAINT UQ__id UNIQUE (_id)
GO

ALTER TABLE sessionStateStore
  ADD CONSTRAINT UQ_SessionStateID UNIQUE (SessionStateID)
GO

ALTER TABLE sessionStateStore
  ADD CONSTRAINT UQ_Code UNIQUE (Code)
GO

ALTER TABLE sessionStateStore
  ADD CONSTRAINT UQ_Name UNIQUE (Name)
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'Session State (Active/Deleted etc.)', 'user', dbo, 'table', 'sessionStateStore'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'sessionStateStore', 'column', '_id'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'sessionStateStore', 'column', '_created'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'sessionStateStore', 'column', '_createdBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'sessionStateStore', 'column', '_createdHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'sessionStateStore', 'column', '_updated'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'sessionStateStore', 'column', '_updatedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'sessionStateStore', 'column', '_updatedHost'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description', 
  'system field', 'user', dbo, 'table', 'sessionStateStore', 'column', '_deleted'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'sessionStateStore', 'column', '_deletedBy'
GO

EXECUTE sys.sp_addextendedproperty 'MS_Description',
  'system field', 'user', dbo, 'table', 'sessionStateStore', 'column', '_deletedHost'
GO