import useDevelopers from '@/hooks/queries/useDeveloper';
import { Stack, Text } from '@mantine/core';
import { useParams } from '@tanstack/react-router';

function DeveloperDetailPage() {
  const id = useParams({
    from: '/developer/$id',
    select: (params) => params.id,
  });

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
