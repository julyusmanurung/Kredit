import {Button, Col, Container, Form, Row} from "react-bootstrap"
import DataTable from 'react-data-table-component'
import DataTableExtensions from 'react-data-table-component-extensions'
import 'react-data-table-component-extensions/dist/index.css'
import React, {useEffect, useState} from "react";
import axios from "axios";
import {rupiahLocale} from "../../utils/utils";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {
    faArrowLeft,
    faArrowRight,
    faBackwardStep,
    faFilter, faFilterCircleXmark,
    faForwardStep
} from "@fortawesome/free-solid-svg-icons";
import './laporan.css'

export const Laporan = () => {
    const localhost = "http://localhost:8080/"
    const paginationRowsPerPageOptions = [5, 15, 20, 25, 30]

    const [laporanData, setLaporanData] = useState({})
    const [branch, setBranch] = useState({})
    const [company, setCompany] = useState("")
    const [alreadyFiltered, setAlreadyFiltered] = useState(false)
    const [selectedBranch, setSelectedBranch] = useState('')
    const [selectedCompany, setSelectedCompany] = useState('')
    const [selectedStartDate, setSelectedStartDate] = useState('')
    const [selectedEndDate, setSelectedEndDate] = useState('')

    useEffect(() => {
        getAllBranch()
        getAllCompany()
        getLaporanData()
        const inputDate = document.querySelectorAll(".input-date")
        const inputSelect = document.querySelectorAll("select")
        for (let i = 0; i < inputDate.length; ++i) {
            localStorage.getItem("mode") === "light" ? inputDate[i].style.colorScheme = "light" : inputDate[i].style.colorScheme = "dark"
        }
        for (let i = 0; i < inputSelect.length; ++i) {
            localStorage.getItem("mode") === "light" ?
                inputSelect[i].classList.add("form-select-light") :
                inputSelect[i].classList.add("form-select-dark")
        }
    }, [])

    const getAllBranch = async() => {
        const res = await axios.get(localhost + "branch")
        setBranch(res.data)
    }

    const getAllCompany =  async() => {
        const res = await axios.get(localhost + "company")
        setCompany(res.data)
    }

    const getLaporanData = async () => {
        const res = await axios.get(localhost +
            "laporan?branch=000" +
            "&channeling_company=" +
            "&start_date=" +
            "&end_date="
        )
        if (res.data.data === null) {
            setLaporanData([])
        } else {
            setLaporanData(res.data)
        }
    }

    const handleFilter = async (e) => {
        e.preventDefault()

        setAlreadyFiltered(true)

        const formData = new FormData(e.currentTarget)
        const body = {
            branch: formData.get("branch"),
            company: formData.get("company"),
            startDate: formData.get("startDate"),
            endDate: formData.get("endDate")
        }

        console.log(body)

        const res = await axios.get(localhost +
            "laporan?branch=" + body.branch +
            "&channeling_company=" + body.company +
            "&start_date=" + body.startDate +
            "&end_date="+ body.endDate
        )
        if (res.data.data === null) {
            setLaporanData([])
        } else {
            setLaporanData(res.data)
        }

    }

    const columns = [
        {
            name: 'No',
            selector: 'rownumber',
        },
        {
            name: 'PPK',
            selector: 'ppk',
            sortable: true,
        },
        {
            name: 'Name',
            selector: 'name',
            sortable: true,
        },
        {
            name: 'CC',
            selector: 'channeling_company',
            sortable: true,
        },
        {
            name: 'DD',
            selector:'drawdown_date',
            format: row => new Date(row.drawdown_date).toLocaleDateString('en-US', {month: '2-digit',day: '2-digit',year: 'numeric'}),
            sortable: true,
        },
        {
            name: 'LA',
            selector: 'loan_amount',
            format: row => rupiahLocale(row.loan_amount),
            sortable: true,
        },
        {
            name: 'LP',
            selector: 'loan_period',
            sortable: true,
        },
        {
            name: 'IE',
            selector: 'interest_effective',
            sortable: true,
        }
    ]

    const data = laporanData.data


    const tableData = {
        columns,
        data,
    };

    const handleResetFilter = () => {
        setAlreadyFiltered(false)
        setSelectedBranch('')
        setSelectedCompany('')
        setSelectedStartDate('')
        setSelectedEndDate('')
        getLaporanData()
    }

    return(
      <Container fluid style={{padding: "1.5rem"}}>
          <Form onSubmit={handleFilter}>
              <Row className="d-flex align-items-center justify-content-left">
                  <h1 className="pb-3">Daftar Approved Data</h1>
                  <Col className="d-flex flex-column gap-2 justify-content-between">
                      <label className="filter-title">Branch: </label>
                      <Form.Select size={2} name="branch" value={selectedBranch} onChange={e => setSelectedBranch(e.target.value)}>
                          <option key={-1} value="000">000 - All Branch</option>
                          {
                              branch?.data? branch.data.map((item,index) => (
                                  <option key={index} value={item.code}>{item.code} - {item.description}</option>)
                              ) : 'Loading...'
                          }
                      </Form.Select>
                  </Col>
                  <Col className="d-flex flex-column align-items-center gap-2 justify-content-start">
                      <label className="filter-title">Company: </label>
                      <Form.Select name="company" value={selectedCompany} onChange={e => setSelectedCompany(e.target.value)}>
                          <option key={-1} value="">All Company</option>
                          {
                              company?.data? company.data.map((item, index) => (
                                  <option key={index+1} value={item.channeling_company}>{item.company_name}</option>)
                              ) : 'Loading...'
                          }
                      </Form.Select>
                  </Col>
                  <Col className="d-flex flex-column align-items-center gap-2 justify-content-start">
                      <label className="filter-title">Start:</label>
                      <Form.Control className="input-date" name="startDate" type="date" value={selectedStartDate} onChange={e => setSelectedStartDate(e.target.value)}/>
                  </Col>
                  <Col className="d-flex flex-column align-items-center gap-2 justify-content-start">
                      <label className="filter-title">End:</label>
                      <Form.Control className="input-date" name="endDate" type="date" value={selectedEndDate} onChange={e => setSelectedEndDate(e.target.value)}/>
                  </Col>
              </Row>
              <div className="pt-4 d-flex gap-3">
                  <Button type="submit" variant="sinarmas"><FontAwesomeIcon icon={faFilter} className="pe-2"/>Filter Data</Button>
                  <Button onClick={() => handleResetFilter()} disabled={alreadyFiltered ? false:true} type="reset" variant="sinarmas2"><FontAwesomeIcon icon={faFilterCircleXmark} className="pe-2" />Reset Filter</Button>
              </div>
          </Form>
          <div className="table-wrapper">
              <DataTableExtensions{...tableData}>
                  <DataTable
                      columns= {columns}
                      data={data}
                      pagination={true}
                      paginationPerPage={5}
                      paginationRowsPerPageOptions={paginationRowsPerPageOptions}
                      responsive={true}
                      paginationIconNext={<FontAwesomeIcon icon={faArrowRight}/>}
                      paginationIconPrevious={<FontAwesomeIcon icon={faArrowLeft}/>}
                      paginationIconFirstPage={<FontAwesomeIcon icon={faBackwardStep}/>}
                      paginationIconLastPage={<FontAwesomeIcon icon={faForwardStep}/>}
                      className="my-5"
                  />
              </DataTableExtensions>
          </div>
      </Container>
  )
}