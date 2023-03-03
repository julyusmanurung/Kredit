import {rupiahLocale} from "../../utils/utils";
import React from "react";

const columns = [
    {
        name: 'No',
        selector: row => row.rownumber,
    },
    {
        name: 'PPK',
        selector: row => row.ppk,
        sortable: true,
    },
    {
        name: 'Name',
        selector: row => row.name,
        sortable: true,
    },
    {
        name: 'Channeling Company',
        selector: row => row.channeling_company,
        sortable: true,
    },
    {
        name: 'Drawdown Date',
        selector: row => row.drawdown_date,
        format: row => new Date(row.drawdown_date).toLocaleDateString('en-US', {month: '2-digit',day: '2-digit',year: 'numeric'}),
        sortable: true,
    },
    {
        name: 'Loan Amount',
        selector: row => row.loan_amount,
        format: row => rupiahLocale(row.loan_amount),
        sortable: true,
    },
    {
        name: 'Loan Period',
        selector: row => row.loan_period,
        sortable: true,
    },
    {
        name: 'Interest Eff',
        selector: row => row.interest_effective,
        sortable: true,
    }
]