import List from "@/components/ui/list";

function UsersList() {
  const users = ['User 1', 'User 2', 'User 3'];
  return (
    <List items={users} />
  );
}

export default UsersList;
