/*
  Warnings:

  - You are about to drop the `_seenBy` table. If the table is not empty, all the data it contains will be lost.
  - Added the required column `password` to the `User` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "_seenBy" DROP CONSTRAINT "_seenBy_A_fkey";

-- DropForeignKey
ALTER TABLE "_seenBy" DROP CONSTRAINT "_seenBy_B_fkey";

-- AlterTable
ALTER TABLE "User" ADD COLUMN     "password" TEXT NOT NULL;

-- DropTable
DROP TABLE "_seenBy";

-- CreateTable
CREATE TABLE "GroupMessageSeenBy" (
    "id" VARCHAR(50) NOT NULL,
    "groupMessageId" VARCHAR(50) NOT NULL,
    "userId" VARCHAR(50) NOT NULL,
    "seenAt" TIMESTAMP(3),

    CONSTRAINT "GroupMessageSeenBy_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "GroupMessageSeenBy_groupMessageId_userId_idx" ON "GroupMessageSeenBy"("groupMessageId", "userId");

-- AddForeignKey
ALTER TABLE "GroupMessageSeenBy" ADD CONSTRAINT "GroupMessageSeenBy_groupMessageId_fkey" FOREIGN KEY ("groupMessageId") REFERENCES "GroupMessage"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "GroupMessageSeenBy" ADD CONSTRAINT "GroupMessageSeenBy_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
