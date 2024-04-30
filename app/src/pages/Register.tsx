import AuthForm from '../components/Auth/AuthForm';

const Register: React.FC = () => {
    return (
        <div className="register-page">
            {/* <h1>Register</h1> */}
            <AuthForm isRegister={true} />
        </div>
    );
}

export default Register;
