import useUser from '@/hooks/queries/useUser';
import { Select, type SelectProps } from '@mantine/core';

function SelectUser({ ...props }: SelectProps) {
  const { listUser } = useUser();

  const queryUser = listUser();
  const dataList = queryUser.data?.users || [];

  return (
    <Select
      {...props}
      data={dataList.map((item) => ({
        label: item.name,
        value: item.id,
      }))}
    />
  );
}

export default SelectUser;
