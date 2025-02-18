# Docs
- GS 001: Requirements based on documented scenarios
- GS 002: Reference Architecture
- GS 003: End-to-end management and orchestration of network slicing
- GR 004: Landscape
- GS 005: Means of Automation
- GS 006: Proof of Concept Framework
- GS 007: Terminology for concepts in ZSM
- GS 008: Cross-domain E2E service lifecycle management
- GS 009-1: Closed-Loop Automation; Part 1: Enablers
- GS 009-2: Closed-Loop Automation; Part 2: Solutions for automation of E2E service and network management use cases
- GR 009-3: Closed-Loop Automation; Part 3: Advanced topics
- GS 010: General Security Aspects
- GR 011: Intent-driven autonomous networks; Generic aspects
- GS 012: Enablers for Artificial Intelligence-based Network and Service Automation
- GS 014: ZSM security aspects
- GR 015: Network Digital Twin
- GS 016: Intent-driven Closed Loops
- GS 018: Network Digital Twin for enhanced zero-touch network and service management

# Scopes
### GS 001: Requirements based on documented scenarios
The present document defines requirements on the zero-touch E2E (End-to-End) network and service management.
Scenarios will be documented and used to derive the requirements.
The requirements will also be considered for the work on the topics Zero-touch network and Service Management
(ZSM) reference architecture [1], ZSM End to end management and orchestration of network slicing [i.7], and ZSM
Inter management domain lifecycle management [i.8].
### GS 002: Reference Architecture
The present document defines and describes the reference architecture for the end-to-end Zero-touch network and
Service Management (ZSM) framework based on a set of user scenarios and requirements documented in ETSI
GS ZSM 001 [i.9].
The reference architecture employs a set of architectural principles, described further in the present document, and a
service-centric architectural model to define at a high level a set of management services for zero-touch network and
service management. It also defines means of management service integration, communication, interoperation, and
organization. Procedures and detailed information models are beyond the scope of the present document.
The reference architecture also defines normative provisions for externally visible management services, defined as part
of the reference architecture, as well as recommendations for their organization. It is assumed that the architectural
patterns introduced in the present document can be used not only for the ZSM framework, but also for architecture and
design of individual management services. 
### GS 003: End-to-end management and orchestration of network slicing
The present document specifies the E2E network slicing management solutions and related management interfaces. The
E2E network slicing including provisioning, performance assurance and fault management of an E2E slice instance
across multiple management domains. 
### GR 004: Landscape
The present document develops a landscape report for Zero-Touch Network and Service Management. It identifies and
includes information about activities in other bodies (such as Standards Developing Organizations, Open Source
Communities, and Industry Associations) that are relevant to the work in ISG ZSM. 
### GS 005: Means of Automation
The present document explores different existing means or approaches to achieve automation. Such means may exist at
different level of management and managed systems, e.g. at service and network management level, at managed
network function level and at managed network level for autonomous optimization. All these means can have a value to
achieve the ZSM goals.
The present document does not address a systematic survey of all standardization activities, open-source or other means
to achieve automation. Rather it evaluates selected existing and proven means of automation at different levels of
managed and management systems, provided by members of the ISG. That comprises for example:
• Alternatives for classic modelling such as intent vs. imperative vs. declarative modelling.
• Lessons learnt from 'model driven automation' of service and policy orchestration, including closed loop.
• Framework for self-managed VNFs based on cloud native network functions and implications.
• Delineate service modelling against 'machine learning'-inspired 'closed loop automation' modelling.
• Autonomous management of networks.
Means of automation may be single elements or composition of elements which enable service and network automation.
They are made visible by this report, and may be considered by the ISG as elements in other work items for
specifications
### GS 006: Proof of Concept Framework
The present document defines a framework to be used by ETSI ISG ZSM to coordinate and promote multi stakeholder
Proofs of Concept (PoC) projects illustrating key aspects of ZSM. Proofs of Concept are an important tool to
demonstrate the viability of a new technology during its early days and or pre-standardization phase.
The main objectives of the ZSM PoC framework are:
• to ensure the PoC projects are scoped around relevant topics for ISG ZSM that require from-the-field input;
• to ensure that the PoC results, lessons learnt and identified gaps are feedback to ISG ZSM;
• to build confidence on the viability of ZSM;
• to encourage the development of a diverse and open ecosystem by fostering the integration of components
from different players;
• to support standardization and industry promotion activities of ISG ZSM.
This framework describes:
• the different roles and responsibilities in the PoC activity process;
• the PoC activity process;
• the acceptance criteria for PoC proposals and reports.
### GS 007: Terminology for concepts in ZSM
The present document provides a glossary of terms and concepts related to Zero touch network and Service
Management (ZSM) with the goal to achieve a common language across all the ETSI ISG ZSM deliverables and to
serve as terminology reference for use across the industry. Where necessary, verbose descriptions providing background
for formal concise definitions will be documented. 
### GS 008: Cross-domain E2E service lifecycle management
The present document investigates the management of End to End (E2E) services across Management Domains (MDs).
It defines the management processes during the lifecycle of E2E services (covering onboarding processes, fulfilment
processes and assurance processes) and describes the interactions between E2E service management domain and
management domains during these processes.
Furthermore, it maps the management services used in the management processes to the northbound interfaces of
selected technology domains and references the underlying specifications of these interfaces. These mappings enable
the automation of lifecycle management across domains. 
### GS 009-1: Closed-Loop Automation; Part 1: Enablers
The present document describes enablers for Closed-Loop Automation (CLA) based on the ZSM architectural
framework. The present document initially specifies Closed Loop-specific requirements and introduces Closed Loops
within the ZSM framework, from both functional and deployment perspectives. The specifications include enablers for
Closed Loop Governance (CLG), covering the definitions of Closed Loop lifecycle management as well as Closed Loop
models. Such enablers allow automatic deployment and configuration of Closed Loops involving both the end-to-end
service management domain and the individual management domains. Then, the present document specifies enablers
for Closed Loops coordination within the ZSM framework, including means for delegation and escalation between
Closed Loops as well as other coordination means. Finally, the deliverable specifies stage-2 generic enablers for
Closed-Loop Automation (CLA) and includes new management services in addition to those identified in ETSI
GS ZSM 002 [2]. Closed loops running entirely within the managed entities are out-of-scope.     
### GS 009-2: Closed-Loop Automation; Part 2: Solutions for automation of E2E service and network management use cases
The present document presents solutions to scenarios related to closed loops using the ZSM specified service
capabilities. New service capabilities are specified where the need arises based on the scenarios solution requirements.
### GR 009-3: Closed-Loop Automation; Part 3: Advanced topics
The present document investigates advanced topics related to closed-loop operations such as learning and cognitive
capabilities (e.g. based on different degrees of use and integration of artificial intelligence technologies), ways to set
and evaluate levels of oversight, autonomy, and operational confidence on the behaviour of the closed loops. The
present document will document problem statements and technical challenges, derive potential requirements, capture,
and evaluate potential solution options, and provide recommendations for further standardization activities.
### GS 010: General Security Aspects
The present document studies the security aspects of the ZSM use cases, framework and solutions, identifies potential
security threats and mitigation considerations to be covered in ZSM standardization activities. It aims to outline a list of
security controls (aka security countermeasures) in order to raise awareness of security aspects that could be considered
in ZSM specifications. The present document will explore the relationship between security controls and
technology-specific solutions. 
### GR 011: Intent-driven autonomous networks; Generic aspects
Intent-based management enables simpler, more user-friendly expressions of input information, and higher flexibility in
automation. Intent is a key enabler to increase automation and make management simpler; therefore the present
document investigates the potential use of intents as key enabler for enhancing autonomous network and service
management within ZSM framework. It provides a formal definition of intents and a list of principles of intent-driven
management, leveraging existing standardization work. Some use cases are also included in the present document to
provide examples of management domains where intents are applicable and capabilities that may be needed. Intentdriven management within the ZSM framework is investigated and the concept of an intent management entity is
introduced, which is responsible for the life cycle management of intents and the exchange of intents between different
management domains. The present document also maps the intent management entity with the concept of closed loops
that is specified in ETSI GS ZSM 009-1 [i.11]. Intent modelling is also investigated, and two different approaches are
proposed. The present document defines intent life cycle phases and a state diagram, together with a set of (mandatory
and optional) interface capabilities that are needed for the life cycle management of intents. Finally, additional aspects
such as conflicts between intents, intent translation, and intent testing are investigated. The present document outlines
potential future work based on the topics explored and the critical areas that were identified in the present document.
### GS 012: Enablers for Artificial Intelligence-based Network and Service Automation
The present document specifies extensions and new capabilities (so-called "AI enablers") for the ZSM framework
reference architecture providing support for the automation of management functionalities and operations based on
Artificial Intelligence (AI), applicable to end-to-end and per management domain. The set of AI-enabling capabilities is
specified as management services, complementing the existing management services defined in ETSI GS ZSM 002 [2].
The focus is on AI-related areas such as data (including data handling and analytics), action, interoperation, governance
and execution environment. Furthermore, the use and integration in the ZSM framework of externally provided AIbased capabilities are taken into account. Security and privacy aspects of AI-enabled network and service automation
are taken into account, where the details would be addressed in a Security related WI.
The present document considers AI-related scenarios defined in ETSI GS ZSM 001 [1], as well as new scenarios, in
order to derive AI-specific requirements. The present document also documents deployment aspects of the above
scenarios to validate the applicability of the AI enablers. Related work from standard development organizations,
open-source projects and other sources are considered and re-used, where applicable, in the development of the
specifications. 
### GS 014: ZSM security aspects
The present document defines the security reference architecture for the Zero-touch network and Service
Management (ZSM) framework based on a set of security capabilities.
The present document specifies a set of security capabilities as management services, complementing the existing
management services defined in ETSI GS ZSM 002 [1], which including adaptive trust relationship between
management domains, dynamic access control and exposure of ZSM service, robustness of AI/ML model, automatic
security governance of management service producer. 
### GR 015: Network Digital Twin
The present document describes the Network Digital Twin concept, investigates its applicability for automation of zerotouch network and service management and introduces existing, emerging and future scenarios that can benefit from it.
Principles and functionality needed to support and utilize the Network Digital Twin for zero-touch network and service
management is introduced, considering also state of the art.
The present document outlines recommendations of additional capabilities needed in the ZSM framework to support
Network Digital Twins.
The present document identifies existing specifications and solutions (both ETSI and external ones) that can be
leveraged to maximize synergies. Collaboration with other SDOs (e.g. in IRTF NMRG, ITU-T SG13) are recommended
when appropriate. 
### GS 016: Intent-driven Closed Loops
The present document specifies capabilities to support the combination of closed-loop automation with intents
originating from ZSM consumers, focusing on intent-driven governance and coordination of closed loops. The present
document includes use cases and additional requirements related to intent-driven aspects of ETSI GS ZSM 009-1 [2].
The present document creates a normative specification covering stages 1 and 2. It also identifies and describes
additions to ETSI GS ZSM 002 [1] and ETSI GS ZSM 009-1 [2], as needed. Related work in ETSI, other SDOs and
open-source projects are considered and used where applicable.
### GS 018: Network Digital Twin for enhanced zero-touch network and service management
The present document specifies extensions and new capabilities to support and integrate digital twin technologies with
the ZSM framework reference architecture in order to enhance end to end zero-touch network and service management
and automation.
The present document defines use cases related to Network Digital Twins (NDTs) to derive specific requirements. It
also documents important NDT principles.
The present normative document is based on the ZSM reference architecture and refers to available standards and open
source works where appropriate. ETSI GR ZSM 015 [i.1] provides background relevant to the present document and
should be read as a companion document. 
