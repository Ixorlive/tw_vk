import { useNavigate } from "react-router-dom";
import { UserType } from "../../entity/UserType";
import "./Header.css";
import MaterialButton from "../MaterialButton/MaterialButton";
import { useEffect, useState } from "react";

export interface HeaderProps {
    user: UserType

    applyTabFilter: any
    applyTimeFilter: any
}

export const Header: React.FC<HeaderProps> = ({ user, applyTabFilter, applyTimeFilter }) => {
    const navigate = useNavigate();
    const [isAuthorized, setIsAuthrized] = useState<boolean>(false);
    const [activeTab, setActiveTab] = useState<string>('all')

    useEffect(() => {
        setIsAuthrized(user.id !== 0)
    }, [user.id]);

    const handleSignOut = () => {
        localStorage.removeItem("token")
        setIsAuthrized(false)
        window.location.reload()
    };

    const handleSwitchTab = (newActiveTab: string) => {
        if (activeTab === newActiveTab) {
            return
        }
        setActiveTab(newActiveTab)
        applyTabFilter(newActiveTab)
    };


    return (
        <header>
            <div className="auth-buttons">
                {
                    isAuthorized === false ? (
                        <>
                            <MaterialButton content="Login" onClick={() => navigate("/login")} />
                            <MaterialButton content="Register" onClick={() => navigate("/register")} />
                        </>
                    ) : (
                        <MaterialButton content="Sign out" onClick={handleSignOut} />
                    )
                }
            </div>
            <div className="filter-tabs">
                <button className={activeTab == 'all' ? "tab active" : "tab"}
                    onClick={() => { handleSwitchTab('all') }}>Everything</button>
                <button className={activeTab == 'my' ? "tab active" : "tab"}
                    onClick={() => { handleSwitchTab('my') }}>My Notes</button>
                <select id="timeFilter" onChange={(e) => {
                    applyTimeFilter(e.target.value === "" ? null : e.target.value)
                }}>
                    <option value="">All</option> // returns 'All'
                    <option value={1}>Last Day</option>
                    <option value={3}>Last Three Days</option>
                    <option value={30}>Last Month</option>
                </select>
            </div>
        </header>
    )
}