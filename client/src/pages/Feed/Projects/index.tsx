import CardMain from '@/components/card-main';
import useProject from '@/hooks/queries/useProject';
import { Stack, Text } from '@mantine/core';

function ProjectsPage() {
  const { listProject } = useProject();
  const queryProject = listProject();
  const dataList = queryProject?.data?.projects || [];

  return (
    <div>
      <Stack>
        {dataList.map((item) => (
          <CardMain key={item.id}>
            <Text>{item.title}</Text>
          </CardMain>
        ))}
      </Stack>
    </div>
  );
}

export default ProjectsPage;
