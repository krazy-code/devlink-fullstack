import { Stack, Text } from '@mantine/core';
import { useParams } from 'react-router';
import useDevelopers from '../../hooks/queries/useDevelopers';

function DeveloperDetailPage() {
  const params = useParams();
  const id = params.id as string;

  const { detailDeveloper } = useDevelopers();

  const queryDetail = detailDeveloper(id);
  const dataDeveloper = queryDetail.data;
  return (
    <Stack>
      <Text>{dataDeveloper?.name}</Text>
    </Stack>
  );
}

export default DeveloperDetailPage;
