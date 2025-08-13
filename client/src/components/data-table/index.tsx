import { Center, LoadingOverlay, Table } from '@mantine/core';
import { randomId } from '@mantine/hooks';
import type React from 'react';
import CardMain from '../card-main';

export interface DataTableColumn<T> {
  accKey: keyof T;
  header: React.ReactNode;
}
export interface DataTableProps<T> {
  data: T[];
  columns: DataTableColumn<T>[];
  isLoading?: boolean;
}

export function DataTable<T>({ data, columns, isLoading }: DataTableProps<T>) {
  return (
    <CardMain>
      <Table align="left">
        <Table.Thead>
          <Table.Tr>
            {columns.map((col) => (
              <Table.Th key={col.accKey as string}>{col.header}</Table.Th>
            ))}
          </Table.Tr>
        </Table.Thead>
        <Table.Tbody>
          {data.map((item) => {
            return (
              <Table.Tr key={randomId()}>
                {columns.map((col) => (
                  <Table.Td key={col.accKey as string} ta="left">
                    {item[col.accKey as keyof typeof item] as React.ReactNode}
                  </Table.Td>
                ))}
              </Table.Tr>
            );
          })}
        </Table.Tbody>
      </Table>
      <Center>
        <LoadingOverlay visible={isLoading} />
      </Center>
    </CardMain>
  );
}
