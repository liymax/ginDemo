import React, {useState} from 'react'
import './App.css'
import {withStyles} from '@material-ui/core/styles'
import Button from '@material-ui/core/Button'
import TablePagination from '@material-ui/core/TablePagination'
import {useCount} from "./hooks/main"
import BusinessTable from "./components/BusinessTable"
import SysMemo from "./components/SysMemo"

const IButton = withStyles({
  root: {textTransform: 'none'}
})(Button);
function App() {
  let [page, setPage] = useState(0);
  function handleChangePage(event, newPage) {
    setPage(newPage);
  }
  let count = useCount(page)
  return (
    <div className="app">
      <header className="app-header">
        <IButton variant="contained" color="primary">Material UI</IButton>
      </header>
      <main className="app-main">
        <BusinessTable pageIndex={page + 1 } />
        <TablePagination
          component="div"
          rowsPerPage={10}
          page={page}
          count={count}
          backIconButtonProps={{
            'aria-label': 'previous page',
          }}
          nextIconButtonProps={{
            'aria-label': 'next page',
          }}
          onChangePage={handleChangePage}
        />
      </main>
      <footer className="app-footer" >
        <SysMemo />
      </footer>
    </div>
  );
}

export default App;
