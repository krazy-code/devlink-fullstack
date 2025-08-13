import type { DataTableColumn } from '@/components/data-table';
import { DataTable } from '@/components/data-table';
import useDeveloper from '@/hooks/queries/useDeveloper';
import type { DeveloperTypeItem } from '@/services/developers/developers.types';
import { Button } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { IconPlus } from '@tabler/icons-react';
import { useMemo } from 'react';
import DeveloperForm from './form';

function AdminDeveloperPage() {
  const [modalFormOpened, modalForm] = useDisclosure();
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
      <Button type="submit" leftSection={<IconPlus />} onClick={modalForm.open}>
        Add Developer
      </Button>

      <DataTable
        data={dataList || []}
        columns={columns}
        isLoading={queryDeveloper.isLoading}
      />
      <DeveloperForm opened={modalFormOpened} onClose={modalForm.close} />
    </>
  );
}

export default AdminDeveloperPage;
