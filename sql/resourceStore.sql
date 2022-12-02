CREATE TABLE resourceStore
(
  _id            int           NOT NULL IDENTITY(1,1),
  resourceID nvarchar(MAX) NOT NULL,
  code           nvarchar(4)   NOT NULL,
  name           nvarchar(MAX) NOT NULL,
  email           nvarchar(MAX) NOT NULL,
  class          nvarchar(MAX) NOT NULL,
  _created       nvarchar(MAX),
  _createdBy     nvarchar(MAX),
  _createdHost   nvarchar(MAX),
  _updated       nvarchar(MAX),
  _updatedBy     nvarchar(MAX),
  _updatedHost   nvarchar(MAX),
  _deleted       nvarchar(MAX),
  _deletedBy     nvarchar(MAX),
  _deletedHost   nvarchar(MAX),
  CONSTRAINT PK_resourceStore PRIMARY KEY (_id)
)
GO