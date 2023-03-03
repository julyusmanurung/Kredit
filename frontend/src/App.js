import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom"
import LoginForm from './components/login/login.component'
import {Home} from "./components/home/home.component"
import Dashboard from "./components/dashboard/dashboard";
import {useEffect, useState} from "react"
import {NavigationBar} from "./components/navbar/navbar.component"
import {SideBar} from "./components/sidebar/sidebar.component"
import {Laporan} from "./components/laporan/laporan";
import {Footer} from "./components/footer/footer";

function App() {
    const [login, setLogin] = useState(false)
    useEffect(() => {
        localStorage.getItem("user") != null ? setLogin(true) : setLogin(false)
    }, [])
    if (!login) {
        return (
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<LoginForm />} />
                </Routes>
            </BrowserRouter>
        )
    } else {
        return (
            <BrowserRouter>
                <div className={localStorage.getItem("mode") === "light" ? "d-flex" : "d-flex dark"}>
                    <SideBar/>
                    <div className="d-flex flex-column right-side">
                        <NavigationBar/>
                        <div className="w-100" style={{overflowX: "hidden"}}>
                            <div className="content">
                                <Routes>
                                    <Route path="/" element={<Dashboard />} />
                                    <Route path="/checklist" element={<Home />} />
                                    <Route path="/laporan" element={<Laporan />} />
                                </Routes>
                            </div>
                            <Footer/>
                        </div>
                    </div>
                </div>
            </BrowserRouter>
        )
    }
}

export default App;
