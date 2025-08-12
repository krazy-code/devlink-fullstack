import type { DataTableColumn } from '@/components/data-table';
import { DataTable } from '@/components/data-table';
import useDeveloper from '@/hooks/queries/useDeveloper';
import type { DeveloperTypeItem } from '@/services/developers/developers.types';
import { useMemo } from 'react';

function AdminDeveloperPage() {
  const { listDeveloper } = useDeveloper();
  const queryDeveloper = listDeveloper();
  const dataList = queryDeveloper.data?.developers || [];
  const columns = useMemo<DataTableColumn<DeveloperTypeItem>[]>(
    () => [
      {
        accKey: 'name',
        header: 'Name',
      },
      {
        accKey: 'github',
        header: 'Github',
      },
    ],
    []
  );
  return (
    <>
      <DataTable
        data={dataList || []}
        columns={columns}
        isLoading={queryDeveloper.isLoading}
      />
    </>
  );
}

export default AdminDeveloperPage;
