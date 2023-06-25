-- CreateTable
CREATE TABLE "Word" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "chatId" BIGINT NOT NULL,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "word" TEXT NOT NULL,
    "translate" TEXT NOT NULL,
    "rememberLevel" INTEGER NOT NULL DEFAULT 0,
    "nextRemindAt" DATETIME NOT NULL,
    "remindedAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "usersChatId" BIGINT,
    CONSTRAINT "Word_usersChatId_fkey" FOREIGN KEY ("usersChatId") REFERENCES "User" ("chatId") ON DELETE SET NULL ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "User" (
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "chatId" BIGINT NOT NULL PRIMARY KEY
);