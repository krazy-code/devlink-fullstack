import CardMain from '@/components/card-main';
import useProject from '@/hooks/queries/useProject';
import { Text, Title } from '@mantine/core';

function ProjectList() {
  const { listProject } = useProject();
  const queryProject = listProject();
  const dataList = queryProject?.data?.projects || [];
  return (
    <CardMain>
      <Title>List Project</Title>
      {dataList.map((item) => (
        <CardMain key={item.id}>
          <Text>{item.title}</Text>
        </CardMain>
      ))}
    </CardMain>
  );
}

export default ProjectList;
