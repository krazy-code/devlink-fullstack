import { Button, LoadingOverlay, SimpleGrid } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { IconPlus } from '@tabler/icons-react';
import useDevelopers from '../../hooks/queries/useDeveloper';
import DeveloperForm from './_containers/developer-form';
import DeveloperItem from './_containers/developer-item';

function DeveloperPage() {
  const [modalFormOpened, modalForm] = useDisclosure();

  const { listDeveloper } = useDevelopers();

  const queryList = listDeveloper();
  const dataList = queryList?.data?.developers || [];
  return (
    <div>
      <LoadingOverlay visible={queryList.isLoading} />
      <SimpleGrid cols={3}>
        {dataList?.map((item) => {
          return <DeveloperItem key={item.id} name={item.name} />;
        })}
      </SimpleGrid>
      <Button type="submit" leftSection={<IconPlus />} onClick={modalForm.open}>
        Add Developer
      </Button>

      <DeveloperForm opened={modalFormOpened} onClose={modalForm.close} />
    </div>
  );
}

export default DeveloperPage;
