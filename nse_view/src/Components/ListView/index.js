import React from "react";
import "./listview.css";

function ListView(props) {
  function GetRows(dataType, data) {
    switch (dataType) {
      case "corpInfo":
        return Object.keys(data).map((key, i) => {
          return (
            <tr key={dataType + i}>
              <td>{key}</td>
              <td>{data[key]}</td>
            </tr>
          );
        });

      case "corpAnnouncement":
        return data.map((item, i) => (
          <tr key={dataType + i}>
            <td>{item}</td>
          </tr>
        ));

      default:
        return (
          <>
            <tr>
              <th>Date</th>
              <th>Purpose</th>
            </tr>
            {data.map((item, i) => (
              <tr key={dataType + i}>
                <td>{item.date}</td>
                <td>{item.purpose}</td>
              </tr>
            ))}
          </>
        );
    }
  }

  function GenTable(dataType, data) {
    return (
      <table className="list-view" style={{ width: "100%" }}>
        <tbody>{GetRows(dataType, data)}</tbody>
      </table>
    );
  }

  return (
    <>
      {props.data && GenTable(props.heading, props.data)}
      {!props.data && <p>nil</p>}
    </>
  );
}

export default ListView;
