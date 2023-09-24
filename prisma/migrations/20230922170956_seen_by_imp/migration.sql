/*
  Warnings:

  - You are about to drop the column `seen` on the `GroupMessage` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "GroupMessage" DROP COLUMN "seen";

-- CreateTable
CREATE TABLE "_seenBy" (
    "A" VARCHAR(50) NOT NULL,
    "B" VARCHAR(50) NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "_seenBy_AB_unique" ON "_seenBy"("A", "B");

-- CreateIndex
CREATE INDEX "_seenBy_B_index" ON "_seenBy"("B");

-- AddForeignKey
ALTER TABLE "_seenBy" ADD CONSTRAINT "_seenBy_A_fkey" FOREIGN KEY ("A") REFERENCES "GroupMessage"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_seenBy" ADD CONSTRAINT "_seenBy_B_fkey" FOREIGN KEY ("B") REFERENCES "User"("id") ON DELETE CASCADE ON UPDATE CASCADE;
