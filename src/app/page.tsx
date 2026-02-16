import { createRoot } from "react-dom"
import "./page.module.css"

const navbar = (
  <nav>
    <h1>Planisphere</h1>
    <ul>
        <li>Menu</li>
        <li>About</li>
        <li>Contact</li>
    </ul>
  </nav>
)

export default function Planisphere() {
	return (
		<div>
			<h1>Planisphere</h1>
			<ul>
				<li>Menu</li>
				<li>About</li>
				<li>Contact</li>
			</ul>
		</div>
	)
}
