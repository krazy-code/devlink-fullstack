import type { DataTableColumn } from '@/components/data-table';
import { DataTable } from '@/components/data-table';
import useUser from '@/hooks/queries/useUser';
import type { UserTypeItem } from '@/services/users/users.type';
import { useMemo } from 'react';

function AdminUserPage() {
  const { listUser } = useUser();
  const queryUser = listUser();
  const dataList = queryUser.data?.users || [];
  const columns = useMemo<DataTableColumn<UserTypeItem>[]>(
    () => [
      {
        accKey: 'name',
        header: 'Name',
      },
      {
        accKey: 'email',
        header: 'Email',
      },
    ],
    []
  );
  return (
    <>
      <DataTable data={dataList || []} columns={columns} />
    </>
  );
}

export default AdminUserPage;
