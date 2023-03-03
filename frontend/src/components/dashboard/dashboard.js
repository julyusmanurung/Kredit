import {Col, Container, Row, Table} from "react-bootstrap"
import axios from "axios"
import React, {useEffect, useState} from "react"
import CustomerIllustration from '../../assets/customer.png'
import StagingIllustration from '../../assets/staging.png'
import ErrorIllustration from '../../assets/error.png'
import './dashboard.css'
import {rupiahLocale} from "../../utils/utils"
import Button from "react-bootstrap/Button";

const Dashboard = () => {
    const localhost = "http://localhost:8080/"

    const [customer, setCustomer] = useState({})
    const [customerLoan, setCustomerLoan] = useState({})
    
    const getCustomerData = async () => {
        const res = await axios.get(localhost + "customer")
        setCustomer(res.data)
    }

    const getTotalApproved = customer?.data?.filter(customer => customer.approval_status === "0").length
    const getTotalData = customer?.data?.length

    const getCustomerLoanData = async () => {
        const res = await axios.get(localhost + "statusapprovalnine")
        setCustomerLoan(res.data)
    }

    const [laporanData, setLaporanData] = useState({})

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

    useEffect(() => {
        getCustomerData()
        getCustomerLoanData()
        getLaporanData()
    }, [])

    return (
        <Container fluid className="cf">
            <h1>Sistem Informasi Pengajuan Kredit <br/><span>Welcome to SIPK Dashboard</span></h1>
            <Row className="dashboard-summaries mx-0">
                <Col className="col-summaries">
                    <img src={CustomerIllustration} alt="customer illustration"></img>
                    <div className="d-flex flex-column">
                        <div className="d-flex align-items-center justify-content-start gap-3">
                            <label className="number">{getTotalData}</label>
                            <Button variant="sinarmas">Detail</Button>
                        </div>
                        <label className="desc">Total Customer</label>
                    </div>
                </Col>
                <Col className="col-summaries">
                    <img src={StagingIllustration} alt="total approved illustration"></img>
                    <div className="d-flex flex-column">
                        <div className="d-flex align-items-center justify-content-start gap-3">
                            <label className="number">{getTotalApproved}</label>
                            <Button variant="sinarmas">Detail</Button>
                        </div>
                        <label className="desc">Approved Customer</label>
                    </div>
                </Col>
                <Col className="col-summaries">
                    <img src={ErrorIllustration} alt="total unapproved illustration"></img>
                    <div className="d-flex flex-column">
                        <div className="d-flex align-items-center justify-content-start gap-3">
                            <label className="number">{getTotalData -  getTotalApproved}</label>
                            <Button variant="sinarmas">Detail</Button>
                        </div>
                        <label className="desc">Unapproved Customer</label>
                    </div>
                </Col>
            </Row>
            <div className="dashboard-section px-5 mt-5">
                <div className="d-flex align-items-center py-3 pe-3 gap-3">
                    <label className="section-title">Top 5 Latest Loan Data</label>
                    <Button variant="sinarmas">Detail</Button>
                </div>
                <Table responsive hover>
                    <thead>
                    <tr>
                        <th className="align-middle">Index</th>
                        <th className="align-middle">PPK</th>
                        <th className="align-middle">Name</th>
                        <th>Channeling <br/>Company</th>
                        <th>Drawdown <br/>Date</th>
                        <th>Loan <br/>Amount</th>
                        <th>Loan <br/>Period</th>
                        <th>Interest <br/>Eff</th>
                    </tr>
                    </thead>
                    <tbody>
                    {
                        customerLoan?.data? customerLoan.data.map((item, index) => (
                            <tr key={index}>
                                <td>{item.rownumber}</td>
                                <td>{item.ppk}</td>
                                <td>{item.name}</td>
                                <td>{item.channeling_company}</td>
                                <td>{new Date(item.drawdown_date).toLocaleDateString('en-US', {
                                    month: '2-digit',day: '2-digit',year: 'numeric'})}</td>
                                <td>{rupiahLocale(item.loan_amount)}</td>
                                <td>{item.loan_period}</td>
                                <td>{item.interest_effective}</td>
                            </tr>)
                        ) : <tr><td className='text-center border' colSpan={9}><b>Tidak ada data</b></td></tr>
                    }
                    </tbody>
                </Table>
            </div>
            <div className="dashboard-section px-5 mt-5">
                <div className="d-flex align-items-center py-3 pe-3 gap-3">
                    <label className="section-title">Top 5 Latest Approved Data</label>
                    <Button variant="sinarmas">Detail</Button>
                </div>
                <Table responsive hover>
                    <thead>
                    <tr>
                        <th className="align-middle">Index</th>
                        <th className="align-middle">PPK</th>
                        <th className="align-middle">Name</th>
                        <th>Channeling <br/>Company</th>
                        <th>Drawdown <br/>Date</th>
                        <th>Loan <br/>Amount</th>
                        <th>Loan <br/>Period</th>
                        <th>Interest <br/>Eff</th>
                    </tr>
                    </thead>
                    <tbody>
                    {
                        laporanData?.data? laporanData.data.map((item, index) => (
                            <tr key={index}>
                                <td>{item.rownumber}</td>
                                <td>{item.ppk}</td>
                                <td>{item.name}</td>
                                <td>{item.channeling_company}</td>
                                <td>{new Date(item.drawdown_date).toLocaleDateString('en-US', {
                                    month: '2-digit',day: '2-digit',year: 'numeric'})}</td>
                                <td>{rupiahLocale(item.loan_amount)}</td>
                                <td>{item.loan_period}</td>
                                <td>{item.interest_effective}</td>
                            </tr>)
                        ) : <tr><td className='text-center border' colSpan={9}><b>Tidak ada data</b></td></tr>
                    }
                    </tbody>
                </Table>
            </div>
        </Container>
    )
}

export default Dashboard