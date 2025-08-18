import CardMain from '@/components/card-main';
import useProject from '@/hooks/queries/useProject';
import { Stack } from '@mantine/core';
import ProjectItem from './project.item';

function ProjectList() {
  const { listProject, likeProject } = useProject();
  const queryProject = listProject();
  const dataList = queryProject?.data?.projects || [];

  return (
    <CardMain>
      <Stack>
        {dataList.map((item) => {
          return (
            <ProjectItem
              key={item.id}
              project={item}
              isLike={
                typeof likeProject.context?.action === 'number'
                  ? likeProject.context.action
                  : 0
              }
              handleClickLike={() =>
                likeProject.mutate({
                  id: item.id,
                  action: likeProject.context?.action === 0 ? 1 : 0,
                })
              }
            />
          );
        })}
      </Stack>
    </CardMain>
  );
}

export default ProjectList;
