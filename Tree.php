<?php

declare(strict_types=1);

final class Tree
{
    /**
     * @param string $directory
     * @param string $prefix
     */
    public function walk(string $directory, string $prefix = ''): void
    {
        $filePaths = \scandir($directory, SCANDIR_SORT_NONE);

        foreach ($filePaths as $index => $filePath) {
            if('.' === $filePath[0]) {
                continue;
            }

            $absolute = \implode(DIRECTORY_SEPARATOR, [$directory, $filePath]);
            $isDir = $this->isDir($absolute);

            if($this->isLastNodeInDirectory($index, $filePaths)) {
                $this->writeln($prefix, '└───', $filePath);
                if(true === $isDir) {
                    $this->walk($absolute, \sprintf("%s\t", $prefix));
                }
             } else {
                $this->writeln($prefix, '├───', $filePath);
                if(true === $isDir) {
                    $this->walk($absolute, \sprintf("%s|\t", $prefix));
                }
            }
        }
    }

    /**
     * @param int $index
     * @param array $filePaths
     * @return bool
     */
    private function isLastNodeInDirectory(int $index, array $filePaths): bool
    {
        return $index === \count($filePaths) - 1;
    }

    /**
     * @param string $prefix
     * @param string $indent
     * @param string $filePath
     */
    private function writeln(string $prefix, string $indent, string $filePath): void
    {
        echo $prefix, $indent, $filePath, PHP_EOL;
    }

    /**
     * @param string $absolute
     * @return bool
     */
    private function isDir(string $absolute): bool
    {
        return true === \is_dir($absolute);
    }
}

$tree = new Tree();
$tree->walk($argv[1]);