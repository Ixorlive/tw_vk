import './MaterialButton.css'

interface MaterialButtonProps {
    id?: string
    content?: string
    onClick?: any
}

const MaterialButton: React.FC<MaterialButtonProps> = ({ id, content, onClick }) => {
    return <button className='button' id={id} onClick={onClick}>{content}</button>
}

export default MaterialButton