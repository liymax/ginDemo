import React from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import {useBusinessList} from '../hooks/main'

function BusinessTable(props) {
  const {pageIndex = 1} = props;
  const rows = useBusinessList(pageIndex);
  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>ID</TableCell>
          <TableCell align="right">Uid</TableCell>
          <TableCell align="right">注册状态</TableCell>
          <TableCell align="right">消费者BG</TableCell>
          <TableCell align="right">更新时间</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {rows.map(row => (
          <TableRow key={row.id}>
            <TableCell component="th" scope="row">
              {row.id}
            </TableCell>
            <TableCell align="right">{row.uid}</TableCell>
            <TableCell align="right">{row.registerStatus}</TableCell>
            <TableCell align="right">{row.customerBg}</TableCell>
            <TableCell align="right">{row.updateAt}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  )
}

export default BusinessTable