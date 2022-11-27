USE [EST]
GO

/****** Object:  Table [dbo].[sessionStore]    Script Date: 27/11/2022 12:11:27 ******/
IF  EXISTS (SELECT *
FROM sys.objects
WHERE object_id = OBJECT_ID(N'[dbo].[sessionStore]') AND type in (N'U'))
DROP TABLE [dbo].[sessionStore]
GO

/****** Object:  Table [dbo].[sessionStore]    Script Date: 27/11/2022 12:11:27 ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[sessionStore]
(
    [_id] [int] NOT NULL IDENTITY(1,1) PRIMARY KEY,
    [Id] [nvarchar](max) NULL,
    [Apptoken] [nvarchar](max) NOT NULL,
    [Createdate] [nvarchar](max) NULL,
    [Createtime] [nvarchar](max) NULL,
    [Uniqueid] [nvarchar](max) NULL,
    [Sessiontoken] [nvarchar](max) NULL,
    [Username] [nvarchar](max) NULL,
    [Password] [nvarchar](max) NULL,
    [Userip] [nvarchar](max) NULL,
    [Userhost] [nvarchar](max) NULL,
    [Appip] [nvarchar](max) NULL,
    [Apphost] [nvarchar](max) NULL,
    [Issued] [nvarchar](max) NULL,
    [Expiry] [nvarchar](max) NULL,
    [Expiryraw] [nvarchar](max) NULL,
    [Expires] [datetime2](7) NULL,
    [Brand] [nvarchar](max) NULL,
    [SessionRole] [nvarchar](max) NULL,
    [_created] [nvarchar](max) NULL,
    [_createdBy] [nvarchar](max) NULL,
    [_createdHost] [nvarchar](max) NULL,
    [_updated] [nvarchar](max) NULL,
    [_updatedBy] [nvarchar](max) NULL,
    [_updatedHost] [nvarchar](max) NULL,
    [_deleted] [nvarchar](max) NULL,
    [_deletedBy] [nvarchar](max) NULL,
    [_deletedHost] [nvarchar](max) NULL,

) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO


