function List(props: { items: string[] }) {
  return (
    <ul>
        {props.items.map((item, index) => (
            <li key={index}>{item}</li>
        ))}
    </ul>
  );
}

export default List;
