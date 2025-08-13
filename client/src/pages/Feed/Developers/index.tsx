import { LoadingOverlay, SimpleGrid } from '@mantine/core';
import useDevelopers from '../../../hooks/queries/useDeveloper';
import DeveloperItem from './_containers/developer-item';

function DeveloperPage() {
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
    </div>
  );
}

export default DeveloperPage;
