# Chemical Engineer Problem Solver

## Project Purpose

**Chemical Engineer Problem Solver** is designed to tackle common chemical engineering problems with ease. The primary objective is to simplify engineering calculations and provide quick solutions for users. This project also serves as a learning opportunity to explore and practice the Go programming language.

### Why This Project?

- **Learning and Growth:** To learn Go programming by implementing real-world engineering challenges.
- **Practical Calculations:** To assist chemical engineers in solving frequent problems quickly and accurately.
- **For Students and Professionals:** To provide a calculation tool that benefits both engineering students and working professionals.

---

## Current Features

- **Chemical Equilibrium Calculations:**
  - Calculates the equilibrium constant (Kc) based on user input.
  - Supports JSON-based input and output for streamlined API interactions.

---

## Future Enhancements

### **Planned Features**

1. **Additional Chemical Engineering Calculations:**

   - Heat transfer calculations.
   - Reaction kinetics simulations.
   - Fluid dynamics computations.

2. **Visualization Tools:**

   - Graphical representation of calculation results.
   - Enhanced data analysis through visual tools.

3. **Learning Mode:**

   - Step-by-step problem-solving guides to educate users.

4. **Database Integration:**
   - Allows users to save and retrieve past calculations for further use.

---

## Technologies and Tools

- **Backend:** Go programming language
- **API:** RESTful endpoints
- **Data Format:** JSON
- **Version Control:** Git and GitHub

---

# Chemical Engineer Problem Solver

A Go-based backend project to solve fundamental chemical and fluid mechanics problems. This project includes handlers, services, and models for solving various engineering-related problems such as equilibrium calculations, mole calculations, and fluid mechanics operations.

---

## **Features**

- **Chemical Equilibrium Calculations**
- **Mole Calculations**
- **Fluid Mechanics Formulas:**
  - Bernoulli's Equation
  - Continuity Equation
  - Darcy-Weisbach Equation
  - Reynolds Number

---

## **Fluid Mechanics Formulas**

### **1. Bernoulli's Equation**

The Bernoulli's equation represents the conservation of energy for a fluid flow:
\[
P_1 + \frac{1}{2} \rho v_1^2 + \rho g h_1 = P_2 + \frac{1}{2} \rho v_2^2 + \rho g h_2
\]

- **Variables:**
  - \(P_1, P_2\): Pressure at points 1 and 2 (Pa)
  - \(\rho\): Fluid density (kg/m³)
  - \(v_1, v_2\): Fluid velocity at points 1 and 2 (m/s)
  - \(h_1, h_2\): Height of the fluid at points 1 and 2 (m)
  - \(g\): Gravitational acceleration (m/s²)

---

### **2. Continuity Equation**

The continuity equation represents the conservation of mass in fluid mechanics:
\[
A_1 v_1 = A_2 v_2
\]

- **Variables:**
  - \(A_1, A_2\): Cross-sectional areas at points 1 and 2 (m²)
  - \(v_1, v_2\): Velocities at points 1 and 2 (m/s)

---

### **3. Darcy-Weisbach Equation**

The Darcy-Weisbach equation calculates the pressure loss in a pipe due to friction:
\[
\Delta P = f \frac{L}{D} \frac{\rho v^2}{2}
\]

- **Variables:**
  - \(\Delta P\): Pressure loss (Pa)
  - \(f\): Friction factor (dimensionless)
  - \(L\): Pipe length (m)
  - \(D\): Pipe diameter (m)
  - \(\rho\): Fluid density (kg/m³)
  - \(v\): Fluid velocity (m/s)

---

### **4. Reynolds Number**

The Reynolds number determines whether the flow is laminar or turbulent:
\[
Re = \frac{\rho v D}{\mu}
\]

- **Variables:**
  - \(Re\): Reynolds number (dimensionless)
  - \(\rho\): Fluid density (kg/m³)
  - \(v\): Fluid velocity (m/s)
  - \(D\): Pipe diameter (m)
  - \(\mu\): Fluid dynamic viscosity (Pa·s)

---

## **How to Use**

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/chemical-engineer-problem-solver.git
   ```

## Installation and Usage

### **Requirements**

- Go 1.21 or later
- Git

### **Setup Steps**

1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/chemical-engineer-problem-solver.git
   cd chemical-engineer-problem-solver
   ```
