import React from "react";
import { ComposableMap, Geographies, Geography } from "react-simple-maps";
import styles from "./Map.module.scss";
const Map = () => {
  return (
    <div className={styles.ComposableMapContainer}>
      <ComposableMap
        projectionConfig={{
          scale: 130,
        }}
        className={styles.ComposableMap}
      >
        <Geographies geography="/features.json">
          {({ geographies }) =>
            geographies.map((geo) => (
              <Geography key={geo.rsmKey} geography={geo} fill={"#142d50"} />
            ))
          }
        </Geographies>
      </ComposableMap>
    </div>
  );
};
// #060f1e
export default Map;
